package repository

type Repository interface {
    Save(original, short string) error

    GetByOriginal(original string) (string, error)
    GetByShort(short string) (string, error)

    ExistsShort(short string) bool
}