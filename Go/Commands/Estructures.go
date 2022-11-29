package Commands

type MBR struct {
	mbr_name_disk      []byte
	mbr_path           []byte
	mbr_size           []byte
	mbr_date           []byte
	mbr_disk_signature []byte
	mbr_fit            []byte
	mbr_particion      [4]Particion
}

type Particion struct {
	part_status []byte
	part_type   []byte
	part_fit    []byte
	part_start  []byte
	part_size   []byte
	part_name   []byte
}
