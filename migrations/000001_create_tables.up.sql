CREATE TABLE houses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES users(id),         -- Uyning egasi (foreign key to users)
    title VARCHAR(255) NOT NULL,                         -- Uyning nomi yoki qisqa tavsifi
    description TEXT,                                    -- Uyning to'liq tavsifi
    latitude FLOAT NOT NULL,                             -- Kenglik koordinatasi
    longitude FLOAT NOT NULL,                            -- Uzunlik koordinatasi
    images_url VARCHAR[] NOT NULL,                       -- Uyning rasmlari (array of URLs)
    total_square_area INT,                               -- Umumiy kvadrat maydoni
    price_per_month DECIMAL(10, 2) NOT NULL,             -- Oylik ijara narxi
    bedrooms INT,                                        -- Yotoq xonalari soni
    bathrooms INT,                                       -- Vannaxonalar soni
    created_at TIMESTAMP DEFAULT NOW(),                  -- Yaratilgan vaqt
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),  -- Yangilangan vaqt
    deleted_at TIMESTAMP                                -- O'chirilgan vaqt (soft delete uchun)
);

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Booking ID
    user_id UUID NOT NULL REFERENCES users(id),          -- User (renter) ID (foreign key to users)
    house_id UUID NOT NULL REFERENCES houses(id),        -- House ID (foreign key to houses)
    start_date DATE NOT NULL,                            -- Ijara boshlanish sanasi
    end_date DATE NOT NULL,                              -- Ijara tugash sanasi
    total_price DECIMAL(10, 2) NOT NULL,                 -- Buyurtma uchun umumiy narx
    status VARCHAR(50) DEFAULT 'pending',                -- Buyurtma holati (pending, confirmed, cancelled)
    created_at TIMESTAMP DEFAULT NOW(),                  -- Buyurtma yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Payment ID
    booking_id UUID NOT NULL REFERENCES bookings(id),    -- Buyurtma ID (foreign key to bookings)
    amount DECIMAL(10, 2) NOT NULL,                      -- To'lov miqdori
    method VARCHAR(50) NOT NULL,                         -- To'lov usuli (Payme, Click, PayPal, etc.)
    status VARCHAR(50) DEFAULT 'pending',                -- To'lov holati (pending, completed, failed)
    created_at TIMESTAMP DEFAULT NOW(),                  -- To'lov yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);

CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Review ID
    house_id UUID NOT NULL REFERENCES houses(id),        -- Uy ID (foreign key to houses)
    user_id UUID NOT NULL REFERENCES users(id),          -- User ID (foreign key to users)
    rating INT CHECK (rating >= 1 AND rating <= 5),      -- Reyting (1-5)
    review_text TEXT,                                    -- Foydalanuvchi sharhi
    created_at TIMESTAMP DEFAULT NOW(),                  -- Sharh yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);
