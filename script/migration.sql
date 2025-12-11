CREATE TABLE products (
    id UUID PRIMARY KEY,
    product_id BIGINT UNIQUE,
    product_code VARCHAR(255) UNIQUE,
    product_name VARCHAR(255),
    product_image VARCHAR(255),
    
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP,
    created_user UUID,
    
    is_favourite BOOLEAN DEFAULT FALSE,
    active BOOLEAN DEFAULT TRUE,
    
    hsn_code VARCHAR(255),
    
    total_stock NUMERIC(20,8) DEFAULT 0
);
CREATE TABLE variants (
    id UUID PRIMARY KEY,
    product_id UUID references products(id),
    name  VARCHAR(255)
);
CREATE TABLE variant_options (
    id UUID PRIMARY KEY,
    variant_id UUID references variants(id),
    value  VARCHAR(255)
);