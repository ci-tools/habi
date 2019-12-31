package cmd

import (
	"habitat/pkg"
	"log"
	"os"

	"habitat/pkg/spec"
)

var Opts struct {
	Dry string `desc:"just render and print specs"`
}

// Cli container
var Cli = pkg.CliCmd{
	Name: "habitat",
	// Opts: Opts,
	Run: func() error {
		specfile, err := os.Open("habitat.yml")
		if err != nil {
			return err
		}
		habi, err := spec.Parse(specfile)
		if err != nil {
			return err
		}
		log.Printf("\n---\n%v\n", habi)
		// pretty.YAML(os.Stdout, habi)
		// for _, res := range habi.Files {
		// 	filestate, err := file.Meta(res.Dst)
		// 	log.Printf("start with file: %s", res.Dst)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	log.Printf("actual file mode: %v", filestate.Mode)
		// 	log.Printf("target file mode: %v", res.Mode)
		// 	log.Printf("actual file sha256: %v", filestate.Sha256)
		// 	log.Printf("target file sha256: %v", res.Sha256)

		// 	if !filestate.Exists {
		// 		log.Printf("File not exists\n")
		// 	}
		// 	if filestate.Exists && filestate.Sha256 != res.Sha256 {
		// 		log.Printf("File exists but hash doesn't match\n")
		// 	}
		// 	if !filestate.Exists || filestate.Sha256 != res.Sha256 {
		// 		// ensure file
		// 		if strings.HasPrefix(res.Src, "https://") {
		// 			log.Printf("getting file...")
		// 			if err := getter.GetFile(res.Dst, res.Src); err != nil {
		// 				return fmt.Errorf("Res err: %w", err)
		// 			}
		// 		} else if res.Src == "" {
		// 			log.Printf("put file content")
		// 			ioutil.WriteFile(res.Dst, []byte(res.Content), res.Mode)
		// 		}
		// 		filestate, err = file.Meta(res.Dst)
		// 		if filestate.Sha256 != res.Sha256 {
		// 			return fmt.Errorf("file hash doesn't match")
		// 		}
		// 	}
		// 	if res.Mode != filestate.Mode {
		// 		log.Printf("File ensure mode")
		// 		if err = os.Chmod(res.Dst, res.Mode); err != nil {
		// 			return err
		// 		}
		// 	}
		// 	log.Printf("end with file: %s", res.Dst)

		// }
		return nil
	}}

// Execute cli
func Execute() error {
	return Cli.Execute()
}
