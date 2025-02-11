package std

import (
	"errors"
	"os"
	"path/filepath"
	"sort"

	"github.com/jkcfg/jk/pkg/__std"

	flatbuffers "github.com/google/flatbuffers/go"
)

// FileInfo returns a response to a FileInfo request, encoded ready to
// send to the V8 worker.
func (r ReadBase) FileInfo(path, module string) []byte {
	base, path, err := r.getPath(path, module)
	if err != nil {
		return fsError(err)
	}
	return fileInfo(base, path)
}

// DirectoryListing returns a response to a Dir request, encoded ready
// to send to the V8 worker.
func (r ReadBase) DirectoryListing(path, module string) []byte {
	base, path, err := r.getPath(path, module)
	if err != nil {
		return fsError(err)
	}
	return directoryListing(base, path)
}

func fileInfo(base, rel string) []byte {
	path := filepath.Join(base, rel)
	info, err := os.Stat(path)
	switch {
	case err != nil:
		return fsError(err)
	case !(info.IsDir() || info.Mode().IsRegular()):
		return fsError(errors.New("not a regular file"))
	}
	return fileInfoResponse(info.Name(), rel, info.IsDir())
}

func fsError(err error) []byte {
	b := flatbuffers.NewBuilder(1024)
	off := stdError(b, err)
	__std.FileSystemResponseStart(b)
	__std.FileSystemResponseAddRetvalType(b, __std.FileSystemRetvalError)
	__std.FileSystemResponseAddRetval(b, off)
	off = __std.FileSystemResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

func fileInfoResponse(name, path string, isdir bool) []byte {
	b := flatbuffers.NewBuilder(1024)
	off := buildFileInfo(b, name, path, isdir)
	__std.FileSystemResponseStart(b)
	__std.FileSystemResponseAddRetvalType(b, __std.FileSystemRetvalFileInfo)
	__std.FileSystemResponseAddRetval(b, off)
	off = __std.FileSystemResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

func buildFileInfo(b *flatbuffers.Builder, name, path string, isdir bool) flatbuffers.UOffsetT {
	nameOff := b.CreateString(name)
	pathOff := b.CreateString(path)
	__std.FileInfoStart(b)
	__std.FileInfoAddName(b, nameOff)
	__std.FileInfoAddPath(b, pathOff)
	if isdir {
		__std.FileInfoAddIsdir(b, 1)
	} else {
		__std.FileInfoAddIsdir(b, 0)
	}
	return __std.FileInfoEnd(b)
}

func directoryListing(base, rel string) []byte {
	path := filepath.Join(base, rel)
	dir, err := os.Open(path)
	if err != nil {
		return fsError(err)
	}
	infos, err := dir.Readdir(0)
	if err != nil {
		return fsError(err)
	}

	// Sort the fileinfos by name, to avoid introducing non-determinism
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name() < infos[j].Name()
	})

	b := flatbuffers.NewBuilder(1024)
	offsets := make([]flatbuffers.UOffsetT, 0, len(infos))
	for i := range infos {
		if infos[i].IsDir() || infos[i].Mode().IsRegular() {
			offsets = append(offsets, buildFileInfo(b, infos[i].Name(), filepath.Join(rel, infos[i].Name()), infos[i].IsDir()))
		}
	}

	__std.DirectoryStartFilesVector(b, len(offsets))
	for i := len(offsets) - 1; i >= 0; i-- {
		b.PrependUOffsetT(offsets[i])
	}
	infoVec := b.EndVector(len(offsets))

	nameOff := b.CreateString(filepath.Base(rel))
	pathOff := b.CreateString(rel)
	__std.DirectoryStart(b)
	__std.DirectoryAddName(b, nameOff)
	__std.DirectoryAddPath(b, pathOff)
	__std.DirectoryAddFiles(b, infoVec)
	dirOff := __std.DirectoryEnd(b)

	__std.FileSystemResponseStart(b)
	__std.FileSystemResponseAddRetvalType(b, __std.FileSystemRetvalDirectory)
	__std.FileSystemResponseAddRetval(b, dirOff)
	off := __std.FileSystemResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}
