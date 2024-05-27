# Kernel Module Packer

The module packer takes a compiled kernel and a list of module binaries, and packs
them into a single binary.

## Building from source
```bash
go build cmd/mp.go
```

## Usage
```bash
./mp -spec <spec-file>
```

## Packing Specification File
The packing specification file is a YAML with the following structure:
```yaml
kernel_path: "kernel.bin"
target_path: "packedKernel.bin"
modules:
  - path: "module1.bin"
    attributes:
      name: "module 1"
  - path: "module2.bin"
    attributes:
      name: "module 2"
      other_attr: "some value"
```

### Fields
- `kernel_path`: specifies where is the kernel binary
- `target_path`: specifies where the packed kernel will be written to
- `modules`: is an array of module specifications. Each contains:
  - `path`: path to the module binary
  - `attributes`: a map containing attributes. Right now it's not being used.