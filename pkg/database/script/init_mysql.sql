
SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `auth_method`;
CREATE TABLE `auth_method` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

INSERT INTO `auth_method` VALUES ('1', '设备密钥');
INSERT INTO `auth_method` VALUES ('2', 'ID²');
INSERT INTO `auth_method` VALUES ('3', 'X.509证书');

DROP TABLE IF EXISTS `custom_topic`;
CREATE TABLE `custom_topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT NULL,
  `permission_id` int(11) DEFAULT NULL,
  `detail` varchar(255) DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `data_format`;
CREATE TABLE `data_format` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

INSERT INTO `data_format` VALUES ('1', 'ICA标准数据格式(Alink JSON)');
INSERT INTO `data_format` VALUES ('2', '透传/自定义');

DROP TABLE IF EXISTS `device`;
CREATE TABLE `device` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT NULL,
  `status_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `device_secret` varchar(255) NOT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `activation_time` datetime DEFAULT NULL,
  `last_on_line_time` datetime DEFAULT NULL,
  `iot_id` varchar(255) NOT NULL,
  `label` varchar(255) DEFAULT NULL,
  `batch_create` tinyint(1) DEFAULT NULL,
  `online` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `device_secret` (`device_secret`),
  UNIQUE KEY `iot_id` (`iot_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `device_status`;
CREATE TABLE `device_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

INSERT INTO `device_status` VALUES ('1', '未激活');
INSERT INTO `device_status` VALUES ('2', '禁用');
INSERT INTO `device_status` VALUES ('3', '在线');
INSERT INTO `device_status` VALUES ('4', '离线');

DROP TABLE IF EXISTS `model`;
CREATE TABLE `model` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mongodb_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `scene` varchar(255) DEFAULT NULL,
  `territory_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `model` VALUES ('1', '1', '车辆定位卡', '公共服务', '1');
INSERT INTO `model` VALUES ('2', '2', '路灯照明', '公共服务', '1');

DROP TABLE IF EXISTS `model_territory`;
CREATE TABLE `model_territory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

INSERT INTO `model_territory` VALUES ('1', '智能城市');
INSERT INTO `model_territory` VALUES ('2', '智能生活');
INSERT INTO `model_territory` VALUES ('3', '智能工业');
INSERT INTO `model_territory` VALUES ('4', '边缘计算');
INSERT INTO `model_territory` VALUES ('5', '商业共享');
INSERT INTO `model_territory` VALUES ('6', '智能模板');
INSERT INTO `model_territory` VALUES ('7', '智能电力');
INSERT INTO `model_territory` VALUES ('8', '智能农业');
INSERT INTO `model_territory` VALUES ('9', '智能建筑');
INSERT INTO `model_territory` VALUES ('10', '智能园区');

DROP TABLE IF EXISTS `mongodb_model`;
CREATE TABLE `mongodb_model` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `concise_model_id` varchar(255) DEFAULT NULL,
  `intact_model_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `network_way`;
CREATE TABLE `network_way` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

INSERT INTO `network_way` VALUES ('1', 'WiFi');
INSERT INTO `network_way` VALUES ('2', '蜂窝(2G/3G/4G/5G)');
INSERT INTO `network_way` VALUES ('3', '以太网');
INSERT INTO `network_way` VALUES ('4', 'LoRaWAN');
INSERT INTO `network_way` VALUES ('5', '其他');

DROP TABLE IF EXISTS `node_type`;
CREATE TABLE `node_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

INSERT INTO `node_type` VALUES ('1', '直连设备');
INSERT INTO `node_type` VALUES ('2', '网关子设备');
INSERT INTO `node_type` VALUES ('3', '网关设备');

DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `category` enum('standard_category','custom_category') DEFAULT 'standard_category',
  `object_model_id` int(11) DEFAULT NULL,
  `mongodb_model_id` int(11) DEFAULT NULL,
  `node_type_id` int(11) DEFAULT NULL,
  `network_way_id` int(11) DEFAULT NULL,
  `data_format_id` int(11) DEFAULT NULL,
  `auth_method_id` int(11) DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `product_key` varchar(255) NOT NULL,
  `product_secret` varchar(255) NOT NULL,
  `status` varchar(255) DEFAULT '开发中',
  `label` varchar(255) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_key` (`product_key`),
  UNIQUE KEY `product_secret` (`product_secret`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `topic_permission`;
CREATE TABLE `topic_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

INSERT INTO `topic_permission` VALUES ('1', '订阅');
INSERT INTO `topic_permission` VALUES ('2', '发布');
INSERT INTO `topic_permission` VALUES ('3', '发布和订阅');

DROP TABLE IF EXISTS `model_function`;
CREATE TABLE `model_function`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `model_id` int(11) NULL DEFAULT NULL,
  `function_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `function_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `data_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `required` tinyint(1) NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

INSERT INTO `model_function` VALUES (1, 1, '属性', '地理位置', 'GeoLocation', 'STRUCT', 1);
INSERT INTO `model_function` VALUES (2, 2, '属性', '工作状态', 'LightStatus', 'BOOL', 1);
INSERT INTO `model_function` VALUES (3, 2, '属性', '调光等级', 'LightAdjustLevel', 'INT', 1);
INSERT INTO `model_function` VALUES (4, 2, '属性', '工作电压', 'LightVolt', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (5, 2, '属性', '工作电流', 'LightCurrent', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (6, 2, '属性', '有功功率值', 'ActivePower', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (7, 2, '属性', '功率因数', 'PowerRatio', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (8, 2, '属性', '用电量', 'PowerConsumption', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (9, 2, '属性', '漏电压', 'DrainVoltage', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (10, 2, '属性', '倾斜角度值', 'TiltValue', 'INT', 1);
INSERT INTO `model_function` VALUES (11, 2, '属性', '故障功率门限', 'ErrorPowerThreshold', 'INT', 1);
INSERT INTO `model_function` VALUES (12, 2, '属性', '故障电流门限', 'ErrorCurrentThreshold', 'FLOAT', 1);
INSERT INTO `model_function` VALUES (13, 2, '属性', '倾斜阈值', 'TiltThreshold', 'INT', 1);
INSERT INTO `model_function` VALUES (14, 2, '属性', '欠压阈值', 'UnderVoltThreshold', 'INT', 1);
INSERT INTO `model_function` VALUES (15, 2, '属性', '过流阈值', 'OverCurrentThreshold', 'INT', 1);
INSERT INTO `model_function` VALUES (16, 2, '属性', '过压阈值', 'OverVoltThreshold', 'INT', 1);
INSERT INTO `model_function` VALUES (17, 2, '属性', '灯具故障使能', 'LightErrorEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (18, 2, '属性', '过流告警使能', 'OverCurrentEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (19, 2, '属性', '过压告警使能', 'OverVoltEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (20, 2, '属性', '欠压告警使能', 'UnderVoltEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (21, 2, '属性', '漏电告警使能', 'LeakageEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (22, 2, '属性', '倾斜告警使能', 'OverTiltEnable', 'BOOL', 1);
INSERT INTO `model_function` VALUES (23, 2, '属性', '灯具故障告警', 'LampError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (24, 2, '属性', '过流告警', 'OverCurrentError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (25, 2, '属性', '过压告警', 'OverVoltError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (26, 2, '属性', '欠压告警', 'UnderVoltError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (27, 2, '属性', '倾斜告警', 'OverTiltError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (28, 2, '属性', '漏电告警', 'LeakageError', 'BOOL', 1);
INSERT INTO `model_function` VALUES (29, 2, '属性', '地理位置', 'GeoLocation', 'STRUCT', 1);
INSERT INTO `model_function` VALUES (30, 2, '属性', '光照值', 'LightLux', 'INT', 0);
INSERT INTO `model_function` VALUES (31, 2, '服务', '设备校时服务', 'TimeReset', '', 0);
INSERT INTO `model_function` VALUES (32, 2, '服务', '时间任务下发', 'SetTimerTask', '', 0);