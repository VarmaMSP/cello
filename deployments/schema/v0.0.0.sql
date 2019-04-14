DROP DATABASE IF EXISTS `phenopod`;
CREATE DATABASE `phenopod`;
USE `phenopod`;

CREATE TABLE `podcast` (
    -- Required Tags
    `Id` INT AUTO_INCREMENT,
    `Title` VARCHAR(255) NOT NULL,
    `Description` TEXT NOT NULL,
    `ImagePath` VARCHAR(500) NOT NULL,
    `Language` VARCHAR(4) NOT NULL,
    `Explicit` TINYINT NOT NULL,

    -- Recommended Tags
    `Author` VARCHAR(255),
    `Link` VARCHAR(500),
    `OwnerName` VARCHAR (255),
    `OwnerEmail` VARCHAR (255),
    
    -- Situational Tags
    `Type` ENUM('episodic', 'serial') DEFAULT 'episodic',
    `Copyright` VARCHAR(255),
    `NewFeedUrl` VARCHAR(500),
    `Block` TINYINT DEFAULT 0,
    `Complete` TINYINT DEFAULT 0,

    -- RSS feed Details
    `FeedUrl` VARCHAR(500) NOT NULL,
    `LastModified` VARCHAR(100),
    `ETag` VARCHAR(255),
    
    -- Episode stats
    `TotalEpisodeCount` INT,
    `LatestEpisodeGuid` VARCHAR(255),
    `LatestEpisodePubDate` DATETIME NOT NULL,

    -- others
    `CreatedAt` DATETIME NOT NULL,
    `UpdatedAt` DATETIME NOT NULL,

    PRIMARY KEY (`Id`),
    UNIQUE KEY (`Title`),
    UNIQUE KEY (`FeedUrl`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `episode` (
    -- Required Tags
    `Id` VARCHAR(255),
    `PodcastId` INT NOT NULL,
    `Title` VARCHAR(255) NOT NULL,
    `AudioUrl` VARCHAR(500) NOT NULL,
    `AudioType` VARCHAR(20) NOT NULL,

    -- Recommended Tags
    `Guid` VARCHAR(255),
    `PubDate` DATETIME,
    `Description` TEXT,
    `Duration` SMALLINT,
    `Link` VARCHAR(500),
    `Explicit` TINYINT,

    -- Situational Tags
    `Episode` SMALLINT,
    `Season` SMALLINT,
    `EpisodeType` ENUM('full', 'trailer', 'bonus') DEFAULT 'full',
    `Block` TINYINT DEFAULT 0,

    -- others
    `CreatedAt` DATETIME NOT NULL,

    PRIMARY KEY (`Id`),
    FOREIGN KEY (`PodcastId`) REFERENCES `podcast` (`Id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


CREATE TABLE `category` (
    `Id` INT,
    `Name` VARCHAR(100) NOT NULL,
    `ParentId` INT, 
    PRIMARY KEY (`Id`),
    FOREIGN KEY (`ParentId`) REFERENCES `category` (`Id`)
        ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `podcast_category` (
    `PodcastId` INT NOT NULL,
    `CategoryId` INT NOT NULL,
    FOREIGN KEY (`PodcastId`) REFERENCES `podcast` (`Id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`CategoryId`) REFERENCES `category` (`Id`)
        ON UPDATE CASCADE ON DELETE NO ACTION
);

INSERT INTO `category` (`Id`, `Name`) VALUES (1, 'Arts');
INSERT INTO `category` (`Id`, `Name`) VALUES (2, 'Business');
INSERT INTO `category` (`Id`, `Name`) VALUES (3, 'Comedy');
INSERT INTO `category` (`Id`, `Name`) VALUES (4, 'Education');
INSERT INTO `category` (`Id`, `Name`) VALUES (5, 'Games & Hobbies');
INSERT INTO `category` (`Id`, `Name`) VALUES (6, 'Government & Organizations');
INSERT INTO `category` (`Id`, `Name`) VALUES (7, 'Health');
INSERT INTO `category` (`Id`, `Name`) VALUES (8, 'Music');
INSERT INTO `category` (`Id`, `Name`) VALUES (9, 'News & Politics');
INSERT INTO `category` (`Id`, `Name`) VALUES (10, 'Religion & Spirituality');
INSERT INTO `category` (`Id`, `Name`) VALUES (11, 'Science & Medicine');
INSERT INTO `category` (`Id`, `Name`) VALUES (12, 'Society & Culture');
INSERT INTO `category` (`Id`, `Name`) VALUES (13, 'Sports & Recreation');
INSERT INTO `category` (`Id`, `Name`) VALUES (14, 'Technology');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (15, 1, 'Design');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (16, 1, 'Fashion & Beauty');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (17, 1, 'Food');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (18, 1, 'Literature');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (19, 1, 'Performing Arts');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (20, 1, 'Visual Arts');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (21, 2, 'Business News');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (22, 2, 'Careers');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (23, 2, 'Investing');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (24, 2, 'Management & Marketing');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (25, 2, 'Shopping');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (26, 4, 'Educational Technology');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (27, 4, 'Higher Education');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (28, 4, 'K-12');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (29, 4, 'Training');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (30, 5, 'Automotive');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (31, 5, 'Aviation');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (32, 5, 'Hobbies');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (33, 5, 'Other Games');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (34, 5, 'Video Games');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (35, 6, 'Local');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (36, 6, 'National');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (37, 6, 'Non-Profit');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (38, 7, 'Alternative Health');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (39, 7, 'Fitness & Nutrition');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (40, 7, 'Self-Help');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (41, 7, 'Sexuality');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (42, 7, 'Kids & Family');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (43, 10, 'Buddhism');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (44, 10, 'Christianity');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (45, 10, 'Hinduism');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (46, 10, 'Islam');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (47, 10, 'Judaism');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (48, 10, 'Other');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (49, 10, 'Spirituality');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (50, 11, 'Medicine');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (51, 11, 'Natural Sciences');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (52, 11, 'Social Sciences');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (53, 12, 'History');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (54, 12, 'Personal Journals');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (55, 12, 'Philosophy');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (56, 12, 'Places & Travel');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (57, 13, 'Amateur');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (58, 13, 'College & High School');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (59, 13, 'Outdoor');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (60, 13, 'Professional');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (61, 13, 'TV & Film');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (62, 14, 'Gadgets');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (63, 14, 'Podcasting');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (64, 14, 'Software How-To');
INSERT INTO `category` (`Id`, `ParentId`, `Name`) VALUES (65, 14, 'Tech News');
