-- +goose Up
-- +goose StatementBegin
INSERT INTO ref_error_codes (code,message,messageID,description,created_date,created_by,updated_date,updated_by) VALUES
	 ('001','Payment is currently being processed','Pembayaran sedang di proses',NULL,'2019-11-21 13:43:22',1,'2019-11-21 09:33:40',1),
	 ('100','Decrypt Failed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:25',1),
	 ('101','Invalid Parent Format',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:26',1),
	 ('102','Invalid Child Format',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:27',1),
	 ('103','Data Expired',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:28',1),
	 ('104','Encrypt Failed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:29',1),
	 ('105','Invalid Hashed Data',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:29',1),
	 ('106','Invalid Encrypted Data',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:30',1),
	 ('108','Invalid Format',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:31',1),
	 ('109','Invalid IP Address',NULL,'ip address not allowed to use this gateway. see ip_whitelist','2019-11-19 13:43:22',1,'2019-11-19 06:43:31',1),
	 ('116','Invalid Access Token',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:32',1),
	 ('117','Incorrect Credential',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 07:24:22',1),
	 ('121','Payment Failed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:33',1),
	 ('122','Invalid Amount','Invalid Amount',NULL,NULL,1,'2020-05-11 04:07:21',1),
	 ('124','Number Only',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:34',1),
	 ('200','Success','Sukses',NULL,'2019-11-19 13:43:22',1,'2020-10-01 21:57:49',1),
	 ('201','Create Success',NULL,NULL,NULL,1,'2021-11-27 13:05:06',1),
	 ('333','Response Data cannot be parsed',NULL,'time limit exceeded to encrypt the data from ecoll','2019-11-19 13:43:22',1,'2019-11-19 06:43:34',1),
	 ('400','Bad Request',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:38',1),
	 ('401','Authentication Failed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:39',1),
	 ('404','Data Not Found',NULL,NULL,'2019-11-19 13:43:22',1,'2022-05-28 17:13:31',1),
	 ('405','Not Allowed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:40',1),
	 ('422','Validation Failed',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:41',1),
	 ('429','Max Request',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:42',1),
	 ('500','Internal Error',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:47',1),
	 ('600','Request Failed','Request Failed',NULL,NULL,1,'2020-05-11 04:07:21',1),
	 ('606','System is under maintenance, please come back later!','Sistem sedang dalam meintenance, silahkan coba kembali dalam beberapa saat.','Maintenance system','2019-11-22 13:13:13',1,'2019-11-22 06:13:13',1),
	 ('665','Invalid Prefix','Invalid Prefix',NULL,NULL,1,'2020-05-08 03:33:18',1),
	 ('667','Invalid Va Type','Invalid Va Type',NULL,NULL,1,'2020-05-11 04:09:02',1),
	 ('901','Invalid Param',NULL,NULL,'2019-11-19 13:43:22',1,'2019-11-19 06:43:51',1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE from `ref_error_codes`;
-- +goose StatementEnd
