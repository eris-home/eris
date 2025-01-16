package storage

import (
	"fmt"
	"os"
	"time"
)

type BackupManager struct {
	BackupDirectory string
}

func NewBackupManager(directory string) *BackupManager {
	return &BackupManager{BackupDirectory: directory}
}

func (b *BackupManager) CreateBackup(filename string, data []byte) error {
	backupPath := fmt.Sprintf("%s/%s_%s.bak", b.BackupDirectory, filename, time.Now().Format("20060102_150405"))
	file, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("failed to create backup file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to backup file: %w", err)
	}

	fmt.Printf("Backup created successfully: %s\n", backupPath)
	return nil
}

func (b *BackupManager) ListBackups() ([]string, error) {
	files, err := os.ReadDir(b.BackupDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to list backups: %w", err)
	}

	var backups []string
	for _, file := range files {
		if !file.IsDir() {
			backups = append(backups, file.Name())
		}
	}

	return backups, nil
}