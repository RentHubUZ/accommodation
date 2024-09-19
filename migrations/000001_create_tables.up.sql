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
CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Booking ID
    user_id UUID NOT NULL REFERENCES users(id),          -- User (renter) ID (foreign key to users)
    house_id UUID NOT NULL REFERENCES houses(id),        -- House ID (foreign key to houses)
    start_date DATE NOT NULL,                            -- Rental start date
    end_date DATE NOT NULL,                              -- Rental end date
    total_price DECIMAL(10, 2) NOT NULL,                 -- Total price for the booking
    status VARCHAR(50) DEFAULT 'pending',                -- Booking status (pending, confirmed, cancelled)
    created_at TIMESTAMP DEFAULT NOW(),                  -- Booking creation time
    updated_at TIMESTAMP DEFAULT NOW()                   -- Last update time
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Payment ID
    booking_id UUID NOT NULL REFERENCES bookings(id),    -- Booking ID (foreign key to bookings)
    amount DECIMAL(10, 2) NOT NULL,                      -- Payment amount
    method VARCHAR(50) NOT NULL,                         -- Payment method (Payme, Click, PayPal, etc.)
    status VARCHAR(50) DEFAULT 'pending',                -- Payment status (pending, completed, failed)
    created_at TIMESTAMP DEFAULT NOW(),                  -- Payment creation time
    updated_at TIMESTAMP DEFAULT NOW()                   -- Last update time
);

CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Review ID
    house_id UUID NOT NULL REFERENCES houses(id),        -- House ID (foreign key to houses)
    user_id UUID NOT NULL REFERENCES users(id),          -- User ID (foreign key to users)
    rating INT CHECK (rating >= 1 AND rating <= 5),      -- Rating (1-5)
    review_text TEXT,                                    -- User's review
    created_at TIMESTAMP DEFAULT NOW(),                  -- Review creation time
    updated_at TIMESTAMP DEFAULT NOW()                   -- Last update time
);
