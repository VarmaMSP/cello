INSERT INTO `task` (`id`, `name`, `type`, `interval_`, `next_run_at`, `active`, `created_at`, `updated_at`) VALUES (8, 'fix_categories', 'IMMEDIATE', 0, 0, 0, 0, 0);

DELETE FROM `podcast_category`;
SET foreign_key_checks = 0;
DELETE FROM `category`;

INSERT INTO `category` (`id`, `name`) VALUES (1, 'Arts');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (2, 1, 'Books');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (3, 1, 'Design');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (4, 1, 'Fashion & Beauty');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (5, 1, 'Food');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (6, 1, 'Performing Arts');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (7, 1, 'Visual Arts');

INSERT INTO `category` (`id`, `name`) VALUES (8, 'Business');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (9, 8, 'Careers');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (10, 8, 'Entrepreneurship');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (11, 8, 'Investing');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (12, 8, 'Management');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (13, 8, 'Marketing');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (14, 8, 'Non-Profit');

INSERT INTO `category` (`id`, `name`) VALUES (15, 'Comedy');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (16, 15, 'Comedy Interviews');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (17, 15, 'Improv');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (18, 15, 'Stand-Up');

INSERT INTO `category` (`id`, `name`) VALUES (19, 'Education');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (20, 19, 'Courses');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (21, 19, 'How To');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (22, 19, 'Language Learning');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (23, 19, 'Self-Improvement');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (111, 19, 'Higher Education');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (112, 19, 'K-12');

INSERT INTO `category` (`id`, `name`) VALUES (24, 'Fiction');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (25, 24, 'Comedy Fiction');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (26, 24, 'Drama');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (27, 24, 'Science Fiction');

INSERT INTO `category` (`id`, `name`) VALUES (28, 'Government');
INSERT INTO `category` (`id`, `name`) VALUES (29, 'History');

INSERT INTO `category` (`id`, `name`) VALUES (30, 'Health & Fitness');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (31, 30, 'Alternative Health');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (32, 30, 'Fitness');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (33, 30, 'Medicine');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (34, 30, 'Mental Health');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (35, 30, 'Nutrition');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (36, 30, 'Sexuality');

INSERT INTO `category` (`id`, `name`) VALUES (37, 'Kids & Family');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (38, 37, 'Education for Kids');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (39, 37, 'Parenting');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (40, 37, 'Pets & Animals');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (41, 37, 'Stories for Kids');

INSERT INTO `category` (`id`, `name`) VALUES (42, 'Leisure');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (43, 42, 'Animation & Manga');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (44, 42, 'Automotive');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (45, 42, 'Aviation');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (46, 42, 'Crafts');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (47, 42, 'Games');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (48, 42, 'Hobbies');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (49, 42, 'Home & Garden');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (50, 42, 'Video Games');

INSERT INTO `category` (`id`, `name`) VALUES (51, 'Music');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (52, 51, 'Music Commentary');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (53, 51, 'Music History');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (54, 51, 'Music Interviews');

INSERT INTO `category` (`id`, `name`) VALUES (55, 'News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (56, 55, 'Business News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (57, 55, 'Daily News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (58, 55, 'Entertainment News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (59, 55, 'News Commentary');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (60, 55, 'Politics');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (61, 55, 'Sports News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (62, 55, 'Tech News');

INSERT INTO `category` (`id`, `name`) VALUES (63, 'Religion & Spirituality');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (64, 63, 'Buddhism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (65, 63, 'Christianity');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (66, 63, 'Hinduism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (67, 63, 'Islam');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (68, 63, 'Judaism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (69, 63, 'Religion');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (70, 63, 'Spirituality');

INSERT INTO `category` (`id`, `name`) VALUES (71, 'Science');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (72, 71, 'Astronomy');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (73, 71, 'Chemistry');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (74, 71, 'Earth Sciences');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (75, 71, 'Life Sciences');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (76, 71, 'Mathematics');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (77, 71, 'Natural Sciences');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (78, 71, 'Nature');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (79, 71, 'Physics');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (80, 71, 'Social Sciences');

INSERT INTO `category` (`id`, `name`) VALUES (81, 'Society & Culture');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (82, 81, 'Documentary');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (83, 81, 'Personal Journals');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (84, 81, 'Philosophy');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (85, 81, 'Places & Travel');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (86, 81, 'Relationships');

INSERT INTO `category` (`id`, `name`) VALUES (87, 'Sports');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (88, 87, 'Baseball');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (89, 87, 'Basketball');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (90, 87, 'Cricket');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (91, 87, 'Fantasy Sports');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (92, 87, 'Football');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (93, 87, 'Golf');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (94, 87, 'Hockey');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (95, 87, 'Rugby');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (96, 87, 'Running');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (97, 87, 'Soccer');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (98, 87, 'Swimming');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (99, 87, 'Tennis');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (100, 87, 'Volleyball');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (101, 87, 'Wilderness');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (102, 87, 'Wrestling');

INSERT INTO `category` (`id`, `name`) VALUES (103, 'Technology');

INSERT INTO `category` (`id`, `name`) VALUES (104, 'True Crime');

INSERT INTO `category` (`id`, `name`) VALUES (105, 'TV & Film');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (106, 105, 'After Shows');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (107, 105, 'Film History');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (108, 105, 'Film Interviews');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (109, 105, 'Film Reviews');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (110, 105, 'TV Reviews');
