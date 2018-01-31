func FormatCSV(records [][]string) (data []byte, err error) {
	buf := new(bytes.Buffer)

	// add utf bom
	buf.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(buf)
	err = w.WriteAll(records)
	if err != nil {
		return
	}

	data = buf.Bytes()
	return
}
