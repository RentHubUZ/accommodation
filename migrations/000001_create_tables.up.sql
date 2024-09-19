CREATE TABLE houses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL,                         -- Uyning egasi (foydalanuvchi) bilan bog'lanadi
    title VARCHAR(255) NOT NULL,                    -- Uyning nomi yoki qisqa tavsifi
    description TEXT,                               -- Uyning to'liq tavsifi
    latitude FLOAT NOT NULL,                        -- coordinates
    longitude FLOAT NOT NULL,                       -- coordinates
    images_url VARCHAR,                             -- uyning rasmlari
    total_square_area INT,                          -- Umumiy kvadrat maydoni
    price_per_month FLOAT NOT NULL,                 -- Oylik ijara narxi
    bedrooms INT,                                   -- Yotoq xonalari soni
    bathrooms INT,                                  -- Vannaxonalar soni
    created_at TIMESTAMP DEFAULT NOW(),             -- Yaratilgan vaqt
    updated_at TIMESTAMP DEFAULT NOW(),             -- Yangilangan vaqt
    deleted_at TIMESTAMP                            -- O'chirilgan vaqti
);