create TABLE IF NOT EXISTS post (
    postID serial primary key ,
    userName varchar,
    Title varchar,
    Description varchar,
    PostTime varchar

);
CREATE TABLE IF NOT EXISTS comments
(
    commentID varchar primary key,
    comment_post_id varchar ,
    CommentedUser varchar ,
    commentTime varchar ,
    CommentDescription varchar ,
    FOREIGN KEY (comment_post_id) REFERENCES post(postID)
   
);
