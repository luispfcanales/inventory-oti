package ports

//personal staff
//nombre-cargo

//descripcion del servicio
type StorageFileService interface {
	Select(string)
}
type FileService interface {
	List(string)
	PreviewRenderPDF() []byte
}
