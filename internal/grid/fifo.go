package grid

import (
    "bufio"
    "encoding/json"
    "os"
)

func ReadFifo(chans *Chans) {
    for {
        f, err := os.OpenFile("/tmp/termgridboard", os.O_RDONLY, 0600)
        if err != nil {
            panic(err)
        }

        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            text := scanner.Text()

            var cell Cell
            if err := json.Unmarshal([]byte(text), &cell); err != nil {
                continue // Skip invalid JSON
            }
            cell.flags = fAll

            //line <- text
            chans.Cell <- cell
        }
        f.Close()
    }
}
