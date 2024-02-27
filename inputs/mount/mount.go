package mount

import (
	"log"
	"os/exec"
	"strings"

	"flashcat.cloud/categraf/config"
	"flashcat.cloud/categraf/inputs"
	"flashcat.cloud/categraf/types"
)

const inputName = "mount"

type MountStats struct {
	config.PluginConfig
	FileSystem string `toml:"filesystem"`
}

func init() {
	inputs.Add(inputName, func() inputs.Input {
		return &MountStats{}
	})
}

func (h *MountStats) Clone() inputs.Input {
	return &MountStats{}
}

func (h *MountStats) Name() string {
	return inputName
}

func (h *MountStats) Gather(slist *types.SampleList) {
	cmd := exec.Command("mount")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("E! failed to gather mount stats:", err)
	}
	mountCount := 0
	mountLines := strings.Split(string(out), "\n")
	for _, line := range mountLines {
		fields := strings.Fields(line)
		if len(fields) >= 6 && fields[0] == h.FileSystem {
			mountCount++
		}
	}
	fields := map[string]interface{}{
		"count": mountCount,
	}
	slist.PushSamples(inputName, fields)
}
