CREATE TABLE `tb_warteg` (
  `warteg_id` varchar(36) NOT NULL,
  `warteg_name` varchar(255) NOT NULL,
  `warteg_desc` varchar(3000) DEFAULT NULL,
  `warteg_addr` varchar(1000) DEFAULT NULL,
  `warteg_contact_name` varchar(50) DEFAULT NULL,
  `warteg_phone` varchar(100) DEFAULT NULL,
  `updated_date` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `is_delete` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`warteg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
