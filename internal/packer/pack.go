package packer

import (
	"errors"
	"io"
	"log"
	"os"
)

func writeKernel(target *os.File, kernelPath string) error {
	log.Println("writing kernel to target...")
	kernel, err := os.Open(kernelPath)
	defer kernel.Close()
	if err != nil {
		return err
	}

	written, err := io.Copy(target, kernel)
	if err != nil {
		return err
	}
	log.Printf("%d bytes written\n", written)
	return nil
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func strtob(str string) []byte {
	strAsBytes := []byte(str)
	strAsBytes = append(strAsBytes, 0)

	return strAsBytes
}

func writeModuleQuantity(target *os.File, quantity int) error {
	log.Println("writing module count to target...")
	written, err := target.Write(i32tob(uint32(quantity)))
	if err != nil {
		return err
	}
	log.Printf("%d bytes written", written)
	return nil
}

func writeModule(target *os.File, spec ModuleSpec) error {
	moduleName, hasNameAttribute := spec.Attributes["name"]
	if !hasNameAttribute {
		return errors.New("module must contain a name attribute")
	}
	log.Printf("reading module %s at %s to memory...", moduleName, spec.Path)
	moduleBinary, err := os.ReadFile(spec.Path)
	if err != nil {
		return err
	}
	log.Printf("%d bytes read to memory", len(moduleBinary))

	log.Print("writing module size to target...")
	written, err := target.Write(i32tob(uint32(len(moduleBinary))))
	if err != nil {
		return err
	}
	log.Printf("%d bytes written\n", written)

	log.Print("writing module name to target...")
	written, err = target.Write(strtob(moduleName))
	if err != nil {
		return err
	}
	log.Printf("%d bytes written\n", written)

	log.Print("writing module to target...")
	written, err = target.Write(moduleBinary)
	log.Printf("%d bytes written\n", written)

	return nil
}

func Pack(spec PackingSpec) error {
	target, err := os.OpenFile(spec.TargetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer target.Close()
	if err != nil {
		return err
	}

	err = writeKernel(target, spec.KernelPath)
	if err != nil {
		return err
	}
	err = writeModuleQuantity(target, len(spec.Modules))
	if err != nil {
		return err
	}

	for _, module := range spec.Modules {
		err = writeModule(target, module)
		if err != nil {
			return err
		}
	}

	log.Println("packing done!")
	return nil
}
