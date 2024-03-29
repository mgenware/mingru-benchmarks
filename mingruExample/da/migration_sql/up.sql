/******************************************************************************************
 * This file was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

CREATE TABLE `departments` (
	`dept_no` CHAR(4) NOT NULL,
	`dept_name` VARCHAR(40) NOT NULL,
	PRIMARY KEY (`dept_no`)
)
CHARACTER SET=utf8mb4
COLLATE=utf8mb4_unicode_ci
;

CREATE TABLE `employees` (
	`emp_no` INT NOT NULL,
	`first_name` VARCHAR(50) NOT NULL,
	`last_name` VARCHAR(50) NOT NULL,
	`gender` VARCHAR(10) NOT NULL,
	`birth_date` DATE NOT NULL,
	`hire_date` DATE NOT NULL,
	PRIMARY KEY (`emp_no`)
)
CHARACTER SET=utf8mb4
COLLATE=utf8mb4_unicode_ci
;

CREATE TABLE `dept_emp` (
	`emp_no` INT NOT NULL,
	`dept_no` CHAR(4) NOT NULL,
	`from_date` DATE NOT NULL,
	`to_date` DATE NOT NULL,
	PRIMARY KEY (`emp_no`, `dept_no`)
)
CHARACTER SET=utf8mb4
COLLATE=utf8mb4_unicode_ci
;
