-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 14 Okt 2020 pada 15.01
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sepatu`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `katalog`
--

CREATE TABLE `katalog` (
  `id` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `merk` varchar(30) NOT NULL,
  `harga` int(10) NOT NULL,
  `model` varchar(50) NOT NULL,
  `gambar` varchar(100) NOT NULL,
  `size` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `katalog`
--

INSERT INTO `katalog` (`id`, `nama`, `merk`, `harga`, `model`, `gambar`, `size`) VALUES
(15, 'Adidas Classic United', 'Adidas', 1200000, 'Sports', '', '45,43,42'),
(16, 'Daun Singkong Tiga', 'Adinda', 450000, 'Sports', 'adinda1', '40,39,38'),
(17, 'Adira Finance', 'Salah', 4500000, 'Casual', 'ASASASASCSC', '40,39,38'),
(18, 'Alphabounce RC', 'Adidas', 1200000, 'Sports', 'ewdwefewfew', '43,42,41'),
(19, 'Mana Bounce 2.0', 'Adidas', 120000, 'Casual', 'dqwdqwdqd', '40,39,38'),
(20, 'TERREX Climacool Voyager Slip-On', 'Adidas', 1200000, 'Casual', 'dqwdqwdqw', '45,43,42'),
(21, 'Crazy Team 2017', 'Adidas', 1200000, 'Sports', 'dqwdwqd', '43,42,41'),
(22, 'Pure Boost', 'Adidas', 1500000, 'Sports', 'ewfewf', '43,42,41'),
(23, 'Specs Swervo Thunder Bolt', 'Specs', 800000, 'Sports', '', '45,43,42');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `katalog`
--
ALTER TABLE `katalog`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `katalog`
--
ALTER TABLE `katalog`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
