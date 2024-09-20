CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    image TEXT DEFAULT 'https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png',
    role VARCHAR(10) CHECK (role IN ('buyer', 'seller', 'agent')), -- enum
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES Users(id),          -- foreign key -> Users.id
    address VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    zip_code VARCHAR(20) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    property_type VARCHAR(50) CHECK (property_type IN ('house', 'apartment', 'condo')), -- enum
    bedrooms INT,
    bathrooms INT,
    square_footage DECIMAL(10, 2),
    year_built INT,
    listing_status VARCHAR(20) CHECK (listing_status IN ('active', 'pending', 'sold')), -- enum
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE IF NOT EXISTS PropertyImages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID NOT NULL REFERENCES Properties(id),   -- foreign key -> Properties.id
    image_url VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id),           -- foreign key -> Users.id
    property_id UUID NOT NULL REFERENCES Properties(id),   -- foreign key -> Properties.id
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL REFERENCES Users(id),         -- foreign key -> Users.id
    receiver_id UUID NOT NULL REFERENCES Users(id),       -- foreign key -> Users.id
    message_content TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id),           -- foreign key -> Users.id
    property_id UUID NOT NULL REFERENCES Properties(id),   -- foreign key -> Properties.id
    rating INT CHECK (rating >= 1 AND rating <= 5),       -- 1-5
    review_text TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE IF NOT EXISTS SearchHistory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id),           -- foreign key -> Users.id
    search_query VARCHAR(255),
    searched_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    buyer_id UUID NOT NULL REFERENCES Users(id),          -- foreign key -> Users.id
    property_id UUID NOT NULL REFERENCES Properties(id),   -- foreign key -> Properties.id
    transaction_date TIMESTAMP DEFAULT NOW(),
    price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('initiated', 'completed', 'canceled')) -- enum
);

CREATE TABLE IF NOT EXISTS Payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES Transactions(id),    -- Booking ID (foreign key to Transactions)
    amount DECIMAL(10, 2) NOT NULL,                      -- To'lov miqdori
    method VARCHAR(50) NOT NULL,                         -- To'lov usuli (Payme, Click, PayPal, etc.)
    status VARCHAR(50) DEFAULT 'pending',                -- To'lov holati (pending, completed, failed)
    created_at TIMESTAMP DEFAULT NOW(),                  -- To'lov yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);

CREATE TABLE IF NOT EXISTS Bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Booking ID
    user_id UUID NOT NULL REFERENCES Users(id),          -- User (renter) ID (foreign key to Users)
    house_id UUID NOT NULL REFERENCES Properties(id),     -- House ID (foreign key to Properties)
    start_date DATE NOT NULL,                            -- Ijara boshlanish sanasi
    end_date DATE NOT NULL,                              -- Ijara tugash sanasi
    total_price DECIMAL(10, 2) NOT NULL,               -- Buyurtma uchun umumiy narx
    status VARCHAR(50) DEFAULT 'pending',                -- Buyurtma holati (pending, confirmed, cancelled)
    created_at TIMESTAMP DEFAULT NOW(),                  -- Buyurtma yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);

CREATE TABLE IF NOT EXISTS Reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    house_id UUID NOT NULL REFERENCES Properties(id),     -- Uy ID (foreign key to Properties)
    user_id UUID NOT NULL REFERENCES Users(id),          -- User ID (foreign key to Users)
    rating INT CHECK (rating >= 1 AND rating <= 5),      -- Reyting (1-5)
    review_text TEXT,                                    -- Foydalanuvchi sharhi
    created_at TIMESTAMP DEFAULT NOW(),                  -- Sharh yaratish vaqti
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()   -- Oxirgi yangilanish vaqti
);
