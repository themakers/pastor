package helm

import (
	"fmt"
	"log/slog"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"helm.sh/helm/v4/pkg/action"
	"helm.sh/helm/v4/pkg/chart/v2/loader"
	"helm.sh/helm/v4/pkg/cli"
	"helm.sh/helm/v4/pkg/repo"

	"github.com/themakers/pastor/lib/archive"
	"github.com/themakers/pastor/lib/http"
	"github.com/themakers/pastor/lib/tmp"
)

// k8s.io/kube-openapi

func RenderChartFromDir(chartDir, chartName, releaseName, namespace string, values map[string]interface{}) string {
	if namespace == "" {
		namespace = "default"
	}

	slog.Info("rendering chart", "dir", chartDir)
	//> Load chart
	ch, err := loader.Load(filepath.Join(chartDir, chartName))
	if err != nil {
		panic(err)
	}

	if releaseName == "" {
		releaseName = ch.Metadata.Name
	}

	{ //> Render
		settings := cli.New()
		actionConfig := new(action.Configuration)
		if err := actionConfig.Init(
			settings.RESTClientGetter(),
			settings.Namespace(),
			"memory",
		); err != nil {
			panic(err)
		}

		install := action.NewInstall(actionConfig)
		install.DryRun = true
		install.ClientOnly = true
		install.ReleaseName = releaseName
		install.Replace = true
		install.Namespace = namespace

		rel, err := install.Run(ch, values)
		if err != nil {
			panic(err)
		} else {
			return rel.Manifest
		}
	}
}

func RenderChartFromArchive(chartArchiveURL, chartName, releaseName, namespace string, values map[string]interface{}) string {
	//var chartName = extractChartNameWithFallback(chartArchiveURL, releaseName)

	//> Download chart archive
	var chartArchiveTmp = tmp.File(chartName)
	//defer chartArchiveTmp.Purge()

	slog.Info("downloading chart archive", "url", chartArchiveURL, "to", chartArchiveTmp.Path)

	http.Download(chartArchiveURL, chartArchiveTmp.Path)

	//> Unpack chart archive
	var chartTmp = tmp.Dir(chartName)
	//defer chartTmp.Purge()

	slog.Info("unpacking chart archive", "from", chartArchiveTmp.Path, "to", chartTmp.Path)
	archive.Untargz(chartArchiveTmp.Path, chartTmp.Path)

	return RenderChartFromDir(chartTmp.Path, chartName, releaseName, namespace, values)
}

func RenderChartFromRemoteRepo(repoURL, chartName, version, releaseName, namespace string, values map[string]interface{}) string {
	//> Download repo index
	var indexTmp = tmp.File("index.yaml")
	//defer indexTmp.Purge()

	slog.Info("downloading repo index", "url", repoURL, "to", indexTmp.Path)

	http.Download(fmt.Sprintf("%s/index.yaml", repoURL), indexTmp.Path)
	idx, err := repo.LoadIndexFile(indexTmp.Path)
	if err != nil {
		panic(err)
	}

	//> Obtain chart archive url
	entry, err := idx.Get(chartName, version)
	if err != nil {
		panic(err)
	}

	slog.Info("found chart", "name", entry.Name, "version", entry.Version, "urls", entry.URLs)

	var chartURL = entry.URLs[0]

	slog.Info("chart url", "url", chartURL)

	return RenderChartFromArchive(chartURL, chartName, releaseName, namespace, values)
}

func extractChartNameWithFallback(rawURL, fallback string) string {
	if name := extractChartName(rawURL); name != "" {
		return name
	} else {
		return fallback
	}
}

func extractChartName(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	filename := path.Base(u.Path) // e.g. "topolvm-15.5.6.tgz"

	lastDash := strings.LastIndex(filename, "-")
	if lastDash == -1 {
		return ""
	}

	return filename[:lastDash]
}
