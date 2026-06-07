package banner

import (
	"bufio"
	"os"
)

func Load(name string) ([]string, error) {
	file, err := os.Open("banners/" + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}