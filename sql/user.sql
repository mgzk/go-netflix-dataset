create user netflix with password 'netflix';

grant select, insert, update, delete on film to netflix;
grant usage on film_id_seq to netflix;
