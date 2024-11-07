
CREATE TABLE customer (id varchar(18) PRIMARY KEY, name varchar(20));

CREATE TABLE deposit (id varchar(18) , denom  varchar(8), amount varchar(40), PRIMARY KEY(id, denom));

CREATE TABLE deposit_history (sn serial PRIMARY KEY, id varchar(18), denom  varchar(8), amount varchar(40), type int8, date DATE, memo varchar(128));
CREATE INDEX idx_id_deposit_history ON deposit_history USING btree(id);

CREATE TABLE send_history (sn serial PRIMARY KEY, sender varchar(18), receiver varchar(18), denom  varchar(8), amount varchar(40), date DATE, memo varchar(128));
CREATE INDEX idx_sender_send_history ON send_history USING btree(sender);
CREATE INDEX idx_receiver_send_history ON send_history USING btree(receiver);
