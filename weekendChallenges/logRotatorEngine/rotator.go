package main

import (
	"compress/gzip"
	_ "compress/gzip" // The blank identifier (_) prevents the "unused import" error until you use it
	"fmt"
	"io"
	_ "io"
	"os"
	_ "os"
	"time"
	_ "time"
)

// CheckAndArchive compresses the log if it exceeds maxSize.
// Returns true if an archive was created, false if skipped.
func CheckAndArchive(filename string, currentSize int64, maxSize int64) bool {

	// TODO 1: Prevent unnecessary rotations
	// If the current log size is less than or equal to maxSize,
	// immediately return false.
	if currentSize <= maxSize {
		return false
	}


	fmt.Printf("\n[ALERT] Log reached %d bytes. Rotating...\n", currentSize)

	// TODO 2: Generate a timestamped archive filename
	//
	// Requirements:
	// 1. Capture the current time.
	// 2. Format it like this: YYYY-MM-DD_HH-MM-SS
	//    (e.g. 2024-06-15_14-30-00)
	// 3. Build a filename matching:
	//    server-YYYY-MM-DD_HH-MM-SS.log.gz
	// 4. Store the final result inside:
	//    archiveName
	formattedTime := time.Now().Format("2006-01-02_15-04-05")
	archiveName := fmt.Sprintf("server-%s.log.gz", formattedTime)
	

	// TODO 3: Open the active log file
	// Open the original log file for reading.
	//
	// If opening fails:
	// - print the error
	// - return false
	inFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("[error] could not open file : %v", err)
		return false
	}


	// TODO 4: Create the archive destination
	// Create the destination .gz archive file.
	//
	// If creation fails:
	// - print the error
	// - close the original log file
	// - return false
	outFile, err := os.Create(archiveName)
	if err != nil {
		fmt.Printf("error  could not create file, %v", err)
		inFile.Close()
		return false
	} 


	// TODO 5: Initialize the gzip writer
	// Wrap the archive file inside gzip.NewWriter().
	//
	// This writer will transparently compress
	// streamed log data.
	gzipWriter := gzip.NewWriter(outFile)



	// TODO 6: Stream and compress log contents
	// Stream data from the original log file
	// into the gzip writer.
	//
	// If compression fails:
	// - print the error
	// - close all open resources safely
	// - return false
	if _, err := io.Copy(gzipWriter, inFile); err != nil {
		fmt.Printf("[error writing data : %v]", err)
		gzipWriter.Close()
		inFile.Close()
		outFile.Close()
		return false;
	}


	// TODO 7: Close resources in the correct order
	//
	// Close:
	// 1. gzip writer
	// 2. archive file
	// 3. original log file
	//
	// Important:
	// The gzip writer MUST close first so the compression footer flushes correctly.
	// If gzWriter.Close() returns an error:
	// - print the error ("[ERROR] Could not finalize gzip:")
	// - close the archive file
	// - close the original log file
	// - return false
	if err := gzipWriter.Close(); err != nil {
		fmt.Printf("[ERROR] Could not finalize gzip: %v", err)
		outFile.Close()
		inFile.Close()
		return false;
	}
	outFile.Close()
	inFile.Close()

	// Uncomment once archiveName exists
	fmt.Println("[ARCHIVE] Saved to:", archiveName)

	// TODO 8: Truncate the active log
	// Reset the original log file back to 0 bytes
	// WITHOUT deleting the file itself.
	//
	// This allows the server process to continue writing into the same file safely.
	//
	// If truncation fails:
	// - print the error ("[ERROR] Could not clear log file:")
	// - return false
	if err := os.Truncate(filename, 0); err != nil {
		fmt.Printf("[ERROR] Could not clear log file : %v", err)
		return false
	}


	return true
}



