CREATE TABLE "PropertyImages" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "property_id" UUID REFERENCES properties(id) NOT NULL,
  "image_url" VARCHAR(255) NOT NULL,
  "uploaded_at" TIMESTAMP DEFAULT (NOW())
);