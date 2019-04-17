package specs

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/replicatedhq/ship/pkg/state"
)

func (r *Resolver) persistToState(root string) error {
	contents := state.UpstreamContents{}

	err := r.FS.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "fs walk")
		}

		// check if this file is a child of `.git`
		// if it is, don't persist it
		if strings.Contains(path, ".git") {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return errors.Wrapf(err, "get relative path to file %s", path)
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		fileContents, err := r.FS.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "read file")
		}

		base64Contents := base64.StdEncoding.EncodeToString(fileContents)

		newFile := state.UpstreamFile{
			FilePath:     relPath,
			FileContents: base64Contents,
		}

		contents.UpstreamFiles = append(contents.UpstreamFiles, newFile)
		return nil
	})

	if err != nil {
		return errors.Wrapf(err, "fetch contents")
	}

	err = r.StateManager.SerializeUpstreamContents(&contents)
	if err != nil {
		return errors.Wrapf(err, "persist contents")
	}

	return nil
}
