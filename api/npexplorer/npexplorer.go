package npexplorer

import (
	"github.com/connectome-neuprint/neuPrintHTTP/api"
	"github.com/connectome-neuprint/neuPrintHTTP/storage"
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	api.RegisterAPI(PREFIX, setupAPI)
}

const PREFIX = "/npexplorer"

type cypherAPI struct {
	Store storage.Store
}

// setupAPI sets up the optionally supported explorer endpoints
func setupAPI(mainapi *api.ConnectomeAPI) error {
	q := &cypherAPI{mainapi.Store}

	endPoint := "findneurons"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getFindNeurons)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getFindNeurons)
	endPoint = "neuronmetavals"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getNeuronMetaVals)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getNeuronMetaVals)
	endPoint = "neuronmeta"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getNeuronMeta)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getNeuronMeta)
	endPoint = "roiconnectivity"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getROIConnectivity)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getROIConnectivity)
	endPoint = "rankedtable"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getRankedTable)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getRankedTable)
	endPoint = "simpleconnections"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getSimpleConnections)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getSimpleConnections)
	endPoint = "roisinneuron"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getROIsInNeuron)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getROIsInNeuron)
	endPoint = "commonconnectivity"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getCommonConnectivity)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getCommonConnectivity)
	endPoint = "autapses"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getAutapses)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getAutapses)
	endPoint = "distribution"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getDistribution)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getDistribution)
	endPoint = "completeness"
	mainapi.SetRoute(api.GET, PREFIX+"/"+endPoint, q.getCompleteness)
	mainapi.SetRoute(api.POST, PREFIX+"/"+endPoint, q.getCompleteness)
	return nil
}

type errorInfo struct {
	Error string `json:"error"`
}

func (ca *cypherAPI) getFindNeurons(c echo.Context) error {
	var reqObject FindNeuronsParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerFindNeurons(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getNeuronMetaVals(c echo.Context) error {
	var reqObject MetaValParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerNeuronMetaVals(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getNeuronMeta(c echo.Context) error {
	var reqObject DatasetParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerNeuronMeta(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getROIConnectivity(c echo.Context) error {
	var reqObject DatasetParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerROIConnectivity(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getRankedTable(c echo.Context) error {
	var reqObject ConnectionsParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerRankedTable(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getSimpleConnections(c echo.Context) error {
	var reqObject ConnectionsParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerSimpleConnections(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getROIsInNeuron(c echo.Context) error {
	var reqObject NeuronNameParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerROIsInNeuron(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getCommonConnectivity(c echo.Context) error {
	var reqObject CommonConnectivityParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerCommonConnectivity(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getAutapses(c echo.Context) error {
	var reqObject DatasetParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerAutapses(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getDistribution(c echo.Context) error {
	var reqObject DistributionParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerDistribution(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

func (ca *cypherAPI) getCompleteness(c echo.Context) error {
	var reqObject CompletenessParams
	c.Bind(&reqObject)
	if data, err := ca.ExplorerCompleteness(reqObject); err != nil {
		errJSON := errorInfo{err.Error()}
		return c.JSON(http.StatusBadRequest, errJSON)
	} else {
		return c.JSON(http.StatusOK, data)
	}
}
