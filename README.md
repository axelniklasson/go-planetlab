# go-planetlab
Go client library for accessing the [PlanetLab API](https://www.planet-lab.org/doc/plc_api). Note that this library only provides wrappers for the methods *not* requiring admin or PI access to the API. Furthermore, no delete methods have been implemented aside from `DeleteKey`, since deletion may have huge consequences.

