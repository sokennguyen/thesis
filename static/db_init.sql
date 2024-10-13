CREATE TABLE test (
    pid INT PRIMARY KEY,
    conversion BOOL,
    bounce BOOL,
    top INT,
    screen VARCHAR(255),
    location VARCHAR(255),
    response_first VARCHAR(255),
    response_second VARCHAR(255),
    likert_1 INT,
    likert_2 INT,
    min_section_id INT,
    min_section_time INT,
    max_section_id INT,
    max_section_time INT,
    min_click_id INT,
    min_click_count INT,
    max_click_id INT,
    max_click_count INT
);

