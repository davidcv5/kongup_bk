package kongdeck

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/davidcv5/kongup/kongfig"
	"github.com/hbagdi/go-kong/kong"
)

func fromKongfig(kongfig *kongfig.Config) (*fileStructure, error) {
	var services []service
	var consumers []consumer
	for _, api := range kongfig.Apis {
		service, err := mapService(&api)
		services = append(services, *service)
		if err != nil {
			return nil, err
		}
	}
	for _, consumer := range kongfig.Consumers {
		c, err := mapConsumer(&consumer)
		consumers = append(consumers, *c)
		if err != nil {
			return nil, err
		}
	}
	var file fileStructure
	file.Services = services
	file.Consumers = consumers
	return &file, nil
}

func mapService(api *kongfig.API) (*service, error) {
	s := &service{}
	s.Name = kong.String(api.Name)
	s.ConnectTimeout = kong.Int(api.Attributes.UpstreamConnectTimeout)
	s.WriteTimeout = kong.Int(api.Attributes.UpstreamSendTimeout)
	s.ReadTimeout = kong.Int(api.Attributes.UpstreamReadTimeout)
	s.Retries = kong.Int(api.Attributes.Retries)
	u, err := url.Parse(api.Attributes.UpstreamURL)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %s", api.Attributes.UpstreamURL)
	}
	s.Protocol = kong.String(u.Scheme)
	s.Host = kong.String(u.Host)
	if len(u.Path) > 0 {
		s.Path = kong.String(u.Path)
	}
	if p, err := strconv.Atoi(u.Port()); err == nil {
		s.Port = kong.Int(p)
	}

	r, err := mapRoute(api)
	if err != nil {
		return nil, err
	}
	s.Routes = []*route{r}

	plugins, err := mapPlugin(&api.Plugins)
	if err != nil {
		return nil, err
	}
	s.Plugins = plugins

	return s, nil
}

func mapRoute(api *kongfig.API) (*route, error) {
	r := &route{}
	r.Name = kong.String(api.Name)
	r.Service = &kong.Service{
		Name: kong.String(api.Name),
	}
	r.PreserveHost = kong.Bool(api.Attributes.PreserveHost)
	r.StripPath = kong.Bool(api.Attributes.StripURI)
	r.Paths = kong.StringSlice(api.Attributes.Uris...)

	return r, nil
}

func mapPlugin(plugins *[]kongfig.Plugin) ([]*plugin, error) {
	result := []*plugin{}
	if plugins == nil || len(*plugins) == 0 {
		return result, nil
	}
	for _, p := range *plugins {
		plugin := &plugin{}
		plugin.Name = kong.String(p.Name)
		plugin.Enabled = kong.Bool(p.Attributes.Enabled)
		p.Attributes.Config.DeepCopyInto(&plugin.Config)
		result = append(result, plugin)
	}
	return result, nil
}

func mapConsumer(con *kongfig.Consumer) (*consumer, error) {
	c := &consumer{}
	c.Username = kong.String(con.Username)
	if con.CustomID != nil {
		c.CustomID = kong.String(*con.CustomID)
	}
	return c, nil
}
