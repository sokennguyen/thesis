CREATE TABLE IF NOT EXISTS main (
    pid INT,
    age INT,
    conversion BOOL,
    bounce BOOL,
    screen_width VARCHAR(255),
    screen_height VARCHAR(255),

    version int,

    click_nav_feat INT DEFAULT 0,
    click_nav_price INT DEFAULT 0,
    click_nav_login INT DEFAULT 0,
    click_nav_start INT DEFAULT 0,
    click_hero_cta INT DEFAULT 0,
    click_hero_login INT DEFAULT 0,
    click_small_feat1_pic INT DEFAULT 0,
    click_small_feat2_pic INT DEFAULT 0,
    click_small_feat3_pic INT DEFAULT 0,
    click_headstart INT DEFAULT 0,
    click_consistency INT DEFAULT 0,
    click_determination INT DEFAULT 0,
    click_big_feat1_img INT DEFAULT 0,
    click_big_feat2_img INT DEFAULT 0,
    click_bigfeat2_cta INT DEFAULT 0,
    click_big_feat3_img INT DEFAULT 0,
    click_bigfeat3_more INT DEFAULT 0,
    click_big_feat4_img INT DEFAULT 0,
    click_bigfeat4_more INT DEFAULT 0,
    click_ending_cta_btn INT DEFAULT 0,

    hover_nav_feat FLOAT DEFAULT 0,
    hover_nav_price FLOAT DEFAULT 0,
    hover_nav_login FLOAT DEFAULT 0,
    hover_nav_start FLOAT DEFAULT 0,
    hover_hero_cta FLOAT DEFAULT 0,
    hover_hero_login FLOAT DEFAULT 0,
    hover_small_feat1_pic FLOAT DEFAULT 0,
    hover_small_feat2_pic FLOAT DEFAULT 0,
    hover_small_feat3_pic FLOAT DEFAULT 0,
    hover_headstart FLOAT DEFAULT 0,
    hover_consistency FLOAT DEFAULT 0,
    hover_determination FLOAT DEFAULT 0,
    hover_big_feat1_img FLOAT DEFAULT 0,
    hover_big_feat2_img FLOAT DEFAULT 0,
    hover_big_feat3_img FLOAT DEFAULT 0,
    hover_big_feat4_img FLOAT DEFAULT 0,
    
    top FLOAT DEFAULT 0,
    hover_hero FLOAT DEFAULT 0,
    hover_feat_list FLOAT DEFAULT 0,
    hover_benefit_list FLOAT DEFAULT 0,
    hover_big_feat_1 FLOAT DEFAULT 0,
    hover_big_feat_2 FLOAT DEFAULT 0,
    hover_big_feat_3 FLOAT DEFAULT 0,
    hover_big_feat_4 FLOAT DEFAULT 0,
    hover_head_logo FLOAT DEFAULT 0,
    hover_hero_title FLOAT DEFAULT 0,
    hover_sub_title FLOAT DEFAULT 0,
    hover_headstart_desc FLOAT DEFAULT 0,
    hover_consistency_desc FLOAT DEFAULT 0,
    hover_flexible_desc FLOAT DEFAULT 0,
    hover_determination_desc FLOAT DEFAULT 0,
    hover_big_feat1_desc FLOAT DEFAULT 0,
    hover_big_feat2_desc FLOAT DEFAULT 0,
    hover_bigfeat2_cta FLOAT DEFAULT 0,
    hover_big_feat3_desc FLOAT DEFAULT 0,
    hover_bigfeat3_more FLOAT DEFAULT 0,
    hover_big_feat4_desc FLOAT DEFAULT 0,
    hover_bigfeat4_more FLOAT DEFAULT 0,
    hover_ending_title FLOAT DEFAULT 0,
    hover_ending_subtitle FLOAT DEFAULT 0,
    hover_ending_cta_btn FLOAT DEFAULT 0,
    hover_footer_logo FLOAT DEFAULT 0,
    hover_footer_product FLOAT DEFAULT 0,
    hover_footer_company FLOAT DEFAULT 0,
    hover_footer_legal FLOAT DEFAULT 0,

    survey_1_1_1 INT DEFAULT 1,
    survey_1_1_2 INT DEFAULT 1,
    survey_1_2_1 INT DEFAULT 1,
    survey_1_2_2 INT DEFAULT 1,
    survey_1_2_3 INT DEFAULT 1,
    survey_1_3_1 INT DEFAULT 1,
    survey_1_3_2 INT DEFAULT 1,
    survey_1_4_1 INT DEFAULT 1,
    survey_1_4_2 INT DEFAULT 1,
    survey_1_4_3 INT DEFAULT 1,
    survey_1_5_1 INT DEFAULT 1,
    survey_1_5_2 VARCHAR(225),
    survey_2_1_1 INT DEFAULT 1,
    survey_2_1_2 INT DEFAULT 1,
    survey_2_2_1 INT DEFAULT 1,
    survey_2_2_2 INT DEFAULT 1,
    survey_2_2_3 INT DEFAULT 1,
    survey_2_3_1 INT DEFAULT 1,
    survey_2_3_2 INT DEFAULT 1,
    survey_2_4_1 INT DEFAULT 1,
    survey_2_4_2 INT DEFAULT 1,
    survey_2_4_3 INT DEFAULT 1,
    survey_2_5_1 INT DEFAULT 1,
    survey_2_5_2 VARCHAR(225),
    survey_6_1 INT DEFAULT 1,
    survey_6_2 VARCHAR(225) 
)
