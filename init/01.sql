CREATE TABLE `log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `request` text COLLATE utf8_unicode_ci,
  `response` text COLLATE utf8_unicode_ci,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci