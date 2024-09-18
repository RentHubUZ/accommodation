CREATE TABLE houses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL,                         -- Uyning egasi (foydalanuvchi) bilan bog'lanadi
    title VARCHAR(255) NOT NULL,                    -- Uyning nomi yoki qisqa tavsifi
    description TEXT,                               -- Uyning to'liq tavsifi
    latitude VARCHAR(255) NOT NULL,                 -- coordinates
    longitude VARCHAR(255) NOT NULL,                -- coordinates
    price_per_month DECIMAL(10, 2) NOT NULL,        -- Oylik ijara narxi
    bedrooms INT,                                   -- Yotoq xonalari soni
    bathrooms INT,                                  -- Vannaxonalar soni
    square_footage DECIMAL(10, 2),                  -- Uyning maydoni
    created_at TIMESTAMP DEFAULT NOW(),             -- Yaratilgan vaqt
    updated_at TIMESTAMP DEFAULT NOW(),             -- Yangilangan vaqt
);
