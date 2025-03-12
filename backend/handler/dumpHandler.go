package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Communinst/GolangWebStore/backend/config"
	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDump(c *gin.Context) {
	dumpCfg := config.MustLoadDumpConfig()
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02_15-04-05")
	filePath := filepath.Join(dumpCfg.Dir, fmt.Sprintf("%s_%s", dumpCfg.Prefix, timeString))

	fmt.Printf("%s\n", filePath)

	cmd := exec.Command("docker", "exec", dumpCfg.ContainerName, "pg_dump", "-U", dumpCfg.Username, "-F", "c", dumpCfg.DbName)
	outputFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DumpServiceInterface.InsertDump(c.Request.Context(), filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dump created successfully"})
}

func (h *Handler) restoreDump(c *gin.Context) {
	var fileName entities.Dump

	if err := c.ShouldBindJSON(&fileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dumpCfg := config.MustLoadDumpConfig()

	// Construct the full path to the source dump file
	sourceDumpPath := filepath.Join(dumpCfg.Dir, fmt.Sprintf("%s_%s", dumpCfg.Prefix, fileName.Filename))

	// Generate target path for restore
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02_15-04-05")
	targetPath := fmt.Sprintf("%s_%s", dumpCfg.RestorePrefix, timeString)

	// Verify the source dump file exists
	if _, err := os.Stat(sourceDumpPath); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Dump file not found at path: %s", sourceDumpPath),
		})
		return
	}

	// Copy dump file to container
	copyCmd := exec.Command("docker", "cp", sourceDumpPath, fmt.Sprintf("%s:%s", dumpCfg.ContainerName, targetPath))
	if copyOutput, err := copyCmd.CombinedOutput(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to copy dump file: %v, output: %s", err, string(copyOutput)),
		})
		return
	}

	// Restore database
	restoreCmd := exec.Command("docker", "exec", dumpCfg.ContainerName, "pg_restore",
		"-U", dumpCfg.Username,
		"--clean",
		"--if-exists",
		"-d", dumpCfg.DbName,
		targetPath)

	if restoreOutput, err := restoreCmd.CombinedOutput(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to restore database: %v, output: %s", err, string(restoreOutput)),
		})
		return
	}

	// Clean up the temporary file in container
	cleanupCmd := exec.Command("docker", "exec", dumpCfg.ContainerName, "rm", targetPath)
	if err := cleanupCmd.Run(); err != nil {
		log.Printf("Warning: Failed to cleanup temporary file: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dump restored successfully"})
}

func (h *Handler) getAllDumps(c *gin.Context) {
	dumps, err := h.service.DumpServiceInterface.GetAllDumps(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dumps)
}
