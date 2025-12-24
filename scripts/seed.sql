-- Seed Data for Procurement System
-- Run this after tables are created by the application

-- =====================================================
-- 1. SUPPLIERS
-- =====================================================
INSERT INTO suppliers (name, email, address) VALUES
('PT Electronic Supplier', 'contact@electronic.com', 'Jl. Sudirman No. 123, Jakarta'),
('PT Furniture Indonesia', 'info@furniture.co.id', 'Jl. Gatot Subroto No. 456, Bandung'),
('PT Office Supplies', 'sales@officesupplies.com', 'Jl. Thamrin No. 789, Surabaya'),
('CV Fresh Produce', 'orders@freshproduce.com', 'Jl. Ahmad Yani No. 321, Semarang')
ON CONFLICT DO NOTHING;

-- =====================================================
-- 2. ITEMS
-- =====================================================
INSERT INTO items (name) VALUES
('Laptop'),
('Mouse Wireless'),
('Keyboard Mechanical'),
('Monitor 24 inch'),
('Webcam HD'),
('Headset Gaming'),
('Meja Kerja'),
('Kursi Kantor'),
('Lemari Arsip'),
('Kacamata'),
('Jam Tangan'),
('Jam Dinding')
ON CONFLICT DO NOTHING;

-- =====================================================
-- 3. SUPPLIER ITEMS (dengan harga dan stok)
-- =====================================================
-- PT Electronic Supplier items
INSERT INTO supplier_items (supplier_id, item_id, price, stock) VALUES
(1, 1, 8500000, 15),  -- Laptop
(1, 2, 150000, 50),   -- Mouse Wireless
(1, 3, 750000, 30),   -- Keyboard Mechanical
(1, 4, 2500000, 20),  -- Monitor 24 inch
(1, 5, 500000, 25),   -- Webcam HD
(1, 6, 800000, 40)    -- Headset Gaming
ON CONFLICT DO NOTHING;

-- PT Furniture Indonesia items
INSERT INTO supplier_items (supplier_id, item_id, price, stock) VALUES
(2, 7, 1500000, 10),  -- Meja Kerja
(2, 8, 2000000, 8),   -- Kursi Kantor
(2, 9, 3500000, 5)    -- Lemari Arsip
ON CONFLICT DO NOTHING;

-- PT Office Supplies items
INSERT INTO supplier_items (supplier_id, item_id, price, stock) VALUES
(3, 2, 120000, 100),  -- Mouse Wireless (harga berbeda)
(3, 3, 700000, 50),   -- Keyboard Mechanical (harga berbeda)
(3, 10, 250000, 60),  -- Kacamata
(3, 11, 500000, 45),  -- Jam Tangan
(3, 12, 150000, 80)   -- Jam Dinding
ON CONFLICT DO NOTHING;

-- CV Fresh Produce items
INSERT INTO supplier_items (supplier_id, item_id, price, stock) VALUES
(4, 10, 200000, 100), -- Kacamata
(4, 11, 450000, 75),  -- Jam Tangan
(4, 12, 120000, 120)  -- Jam Dinding
ON CONFLICT DO NOTHING;

-- =====================================================
-- SUMMARY
-- =====================================================
-- Users: 3 (1 admin, 2 users) - Password perlu di-hash manual atau register via API
-- Suppliers: 4
-- Items: 12
-- Supplier Items: 18 (kombinasi supplier-item dengan harga dan stok berbeda)
-- 
-- Note: 
-- - Beberapa item dijual oleh multiple suppliers dengan harga berbeda
-- - Stock sudah diisi untuk simulasi pembelian
-- - Password users harus di-register manual via API endpoint
-- =====================================================

-- Query untuk verifikasi data:
-- SELECT COUNT(*) as total_suppliers FROM suppliers;
-- SELECT COUNT(*) as total_items FROM items;
-- SELECT COUNT(*) as total_supplier_items FROM supplier_items;
-- 
-- Query untuk melihat inventory lengkap:
-- SELECT 
--   i.name as item_name,
--   s.name as supplier_name,
--   si.price,
--   si.stock
-- FROM supplier_items si
-- JOIN items i ON si.item_id = i.id
-- JOIN suppliers s ON si.supplier_id = s.id
-- ORDER BY s.name, i.name;
