package iamutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/a8m/envsubst"

	"gopkg.in/yaml.v2"
)

// ServiceWrapper ...
type ServiceWrapper interface {
	GetName() string
	Init(DockerComposePod, IamServiceSetting) error
	GetSettings() IamServiceSetting
	GetActive() (string, string, bool)
	GetPod() DockerComposePod
	GetMaster() DockerComposeContainer
	GetSidecars() map[string]DockerComposeContainer
	GetAllContainers() map[string]DockerComposeContainer
}

// IamConfigYaml ...
type IamConfigYaml struct {
	path               string
	Debug              bool                         `yaml:"debug"`
	IamDir             string                       `yaml:"iamDir"`
	Registries         map[string]string            `yaml:"registries"`
	IamServiceSettings map[string]IamServiceSetting `yaml:"services"`
	ServicesToIgnore   []string                     `yaml:"servicestoignore"`
}

// InitFromFile ..
func (i *IamConfigYaml) InitFromFile(configpath string) error {
	userpath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	i.path = userpath + configpath

	err = yaml.Unmarshal(i.GetSourceYamlBytes(), i)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range i.IamServiceSettings {
		v.Name = k
		i.IamServiceSettings[k] = v
	}
	return nil
}

// Update ...
func (i *IamConfigYaml) Update() error {
	newyaml, err := yaml.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(i.path, newyaml, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// PrintSourceYaml ...
func (i *IamConfigYaml) PrintSourceYaml() error {
	fmt.Println(i.GetSourceYaml())
	return nil
}

// GetSourceYaml ...
func (i *IamConfigYaml) GetSourceYaml() string {
	return string(i.GetSourceYamlBytes())
}

// GetSourceYamlBytes ...
func (i *IamConfigYaml) GetSourceYamlBytes() []byte {
	iamconf, err := ioutil.ReadFile(i.path)
	if err != nil {
		log.Fatal(err)
	}
	return iamconf
}

// GetDockerComposePods ...
func (i *IamConfigYaml) GetDockerComposePods() map[string]DockerComposePod {
	files, err := ioutil.ReadDir(i.IamDir)
	if err != nil {
		log.Fatal(err)
	}

	pods := map[string]DockerComposePod{}

	for _, f := range files {
		// ignore directories
		if !f.IsDir() {
			continue
		}
		// ignore hidden files
		if f.Name()[:1] == "." {
			continue
		}

		// ugly way to check if folder is on ignore list
		skip := false
		for _, k := range i.ServicesToIgnore {
			if f.Name() == k {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		iamServiceDir := i.IamDir + "/" + f.Name()
		subfiles, err := ioutil.ReadDir(iamServiceDir)
		if err != nil {
			log.Fatal(err)
		}
		// check files in service's directory
		for _, sf := range subfiles {
			if sf.Name() == "docker-compose.yml" {
				absoluteFileName := iamServiceDir + "/" + sf.Name()
				pod := DockerComposePod{}
				pod.InitFromFile(absoluteFileName)
				pod.DefineSidecars(f.Name())
				pods[f.Name()] = pod
			}
		}
	}
	return pods
}

// IamServiceSetting ...
type IamServiceSetting struct {
	Name   string
	Active bool   `yaml:"active"`
	Env    string `yaml:"env"`
}

// ToggleActive ...
func (i *IamServiceSetting) ToggleActive() error {
	i.Active = !i.Active
	return nil
}

// GetEnvAsEmoji ...
func (i *IamServiceSetting) GetEnvAsEmoji() string {
	if i.Env == "prod" {
		return "🐳"
	}
	return "🚧"
}

// DockerComposePod ...
type DockerComposePod struct {
	Path                   string
	Version                string                            `yaml:"version"`
	DockerComposeContainer map[string]DockerComposeContainer `yaml:"services"`
	// TODO volume and network and ??? + preserve comments?
}

// InitFromFile ...
func (d *DockerComposePod) InitFromFile(path string) {
	d.Path = path
	dockerComposeYamlContent, err := envsubst.ReadFile(path)
	if err != nil {
		log.Fatalf("envsubst error: %v", err)
	}

	err = yaml.Unmarshal(dockerComposeYamlContent, d)
	if err != nil {
		log.Println(err.Error())
	}

	for k, v := range d.DockerComposeContainer {
		v.Name = k
		d.DockerComposeContainer[k] = v
	}
}

// Save ...
func (d *DockerComposePod) Save() error {
	if d.Path == "" {
		// TODO return error
		return nil
	}
	return nil
	// TODO maybe the docker compose file should not be manipulated ??? => work with env vars
	// dockerComposeYamlContent, err := envsubst.ReadFile(path)
	// if err != nil {
	// 	log.Fatalf("envsubst error: %v", err)
	// }

	// err = yaml.Unmarshal(dockerComposeYamlContent, d)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for k,v := range d.DockerComposeContainer {
	// 	v.Name = k
	// }
}

// DefineSidecars walk through DockerComposeContainers and find the master (decalre others as sidecars)
func (d *DockerComposePod) DefineSidecars(servicename string) {
	master := &DockerComposeContainer{}
	for containerName, cont := range d.DockerComposeContainer {
		cont.Name = containerName
		if containerName == servicename {
			cont.DefineSidecar(nil)
			d.DockerComposeContainer[containerName] = cont
			master = &cont
		}
	}
	for containerName, cont := range d.DockerComposeContainer {
		cont.Name = containerName
		if containerName != servicename {
			cont.DefineSidecar(master)
			d.DockerComposeContainer[containerName] = cont
		}
	}
}

// GetMaster ...
func (d *DockerComposePod) GetMaster() DockerComposeContainer {
	for _, cont := range d.DockerComposeContainer {
		if cont.sidecarOf == nil {
			return cont
		}
	}
	// TODO: throw error
	return DockerComposeContainer{}
}

// GetSidecars ...
func (d *DockerComposePod) GetSidecars() map[string]DockerComposeContainer {
	sidecars := map[string]DockerComposeContainer{}
	for name, cont := range d.DockerComposeContainer {
		if cont.sidecarOf != nil {
			sidecars[name] = cont
		}
	}
	return sidecars

}

// DockerComposeContainer ...
type DockerComposeContainer struct {
	Name        string
	sidecarOf   *DockerComposeContainer
	Volumes     []string                       `yaml:"volumes"`
	Healthcheck map[string]interface{}         `yaml:"healthcheck"`
	Restart     string                         `yaml:"restart"`
	EnvFile     []string                       `yaml:"env_file"`
	Image       string                         `yaml:"image"`
	Ports       []string                       `yaml:"ports"`
	EnvVars     map[string]interface{}         `yaml:"environment"`
	Networks    map[string]map[string][]string `yaml:"networks"`
	WorkDir     string                         `yaml:"working_dir"`
}

// DefineSidecar ...
func (d *DockerComposeContainer) DefineSidecar(mastercontainer *DockerComposeContainer) {
	d.sidecarOf = mastercontainer
}

// IsSidecar ...
func (d *DockerComposeContainer) IsSidecar() bool {
	return d.sidecarOf != nil
}

// SidecarOf ...
func (d *DockerComposeContainer) SidecarOf() *DockerComposeContainer {
	return d.sidecarOf
}

// WrapServices ...
func (i *IamConfigYaml) WrapServices(sw ServiceWrapper, service string) {
	fmt.Println(i.GetDockerComposePods())
	pod := i.GetDockerComposePods()[service]
	fmt.Println("Name pod in utils.go: " + pod.GetMaster().Name)

	sw.Init(pod, i.IamServiceSettings[service])
}

// CliService ...
type CliService struct {
	Name    string
	Path    string
	Setting IamServiceSetting
	Pod     DockerComposePod
}

// PrintSimple ...
func (cs *CliService) PrintSimple() {
	fmt.Println(cs.Name)
}

// GetActive ...
func (cs *CliService) GetActive() (string, string, bool) {
	if cs.Setting.Active {
		return "active", "act", true
	}
	return "", "", false
}

// GetName ...
func (cs *CliService) GetName() string {
	return cs.Name
}

// Init ...
func (cs *CliService) Init(pod DockerComposePod, settings IamServiceSetting) error {

	cs.Pod = pod
	cs.Path = filepath.Join(pod.Path, "..")
	cs.Name = pod.GetMaster().Name
	cs.Setting = settings
	return nil
}

// GetSettings ...
func (cs *CliService) GetSettings() IamServiceSetting {
	return cs.Setting
}

// GetPod ...
func (cs *CliService) GetPod() DockerComposePod {
	return cs.Pod
}

// GetMaster ...
func (cs *CliService) GetMaster() DockerComposeContainer {
	return cs.Pod.GetMaster()
}

// GetSidecars ...
func (cs *CliService) GetSidecars() map[string]DockerComposeContainer {
	if cs.Pod.GetSidecars() == nil {
		return nil
	}
	return cs.Pod.GetSidecars()
}

// GetAllContainers ...
func (cs *CliService) GetAllContainers() map[string]DockerComposeContainer {
	return cs.Pod.DockerComposeContainer
}

// GenerateCLIServices ...
func GenerateCLIServices(conf IamConfigYaml) map[string]CliService {
	cliServices := map[string]CliService{}
	for k, v := range conf.GetDockerComposePods() {
		a := *CreateCLIService()
		a.Init(v, conf.IamServiceSettings[k])
		cliServices[a.GetName()] = a
	}
	return cliServices
}

// CreateCLIService ...
func CreateCLIService() *CliService {
	return &CliService{}
}
