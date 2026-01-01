# Database Migrations

This directory contains database migration files for production environments.

## GORM Model Compatibility

The migration files are designed to match the GORM model structure:

### GORM Model Fields (gorm.Model)
When you use `gorm.Model`, it automatically includes:
- `ID` (uint, primary key)
- `CreatedAt` (time.Time)
- `UpdatedAt` (time.Time)
- `DeletedAt` (gorm.DeletedAt, for soft deletes)

### Migration Files
The migration files create tables with the same structure:
- `id` (BIGSERIAL PRIMARY KEY) - matches GORM's uint ID
- `created_at` (TIMESTAMP WITH TIME ZONE) - matches CreatedAt
- `updated_at` (TIMESTAMP WITH TIME ZONE) - matches UpdatedAt
- `deleted_at` (TIMESTAMP WITH TIME ZONE) - matches DeletedAt
- Indexes on `deleted_at` for soft delete queries

## Usage

### Development Environment
- AutoMigrate is used automatically (no need to run migrations manually)
- Set `APP_ENV` to empty, "dev", or "development"

### Production Environment
- Migrations are run automatically on application startup
- Set `APP_ENV` to "production" or any other value

### Manual Migration Commands (if needed)
```bash
# Run migrations up
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Rollback migrations
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down

# Check migration version
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" version
```

## Creating New Migrations

When you add new fields to GORM models, create corresponding migration files:

```bash
migrate create -ext sql -dir migrations -seq add_new_field_to_books
```

This will create two files:
- `XXXXXX_add_new_field_to_books.up.sql` - for applying the change
- `XXXXXX_add_new_field_to_books.down.sql` - for reverting the change
