create TABLE IF NOT EXISTS post (
    postID serial primary key ,
    userName varchar,
    Title varchar,
    Description varchar,
    PostTime varchar

);
CREATE TABLE IF NOT EXISTS comments
(
    commentId serial primary key,
    comment_post_id serial ,
    CommentedUser varchar ,
    commentTime varchar ,
    CommentDescription varchar ,
    FOREIGN KEY (comment_post_id) REFERENCES post(postID)
   
);
