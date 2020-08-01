-- https://gist.github.com/colophonemes/9701b906c5be572a40a84b08f4d2fa4e

-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- table
CREATE TABLE IF NOT EXISTS public.inputs
(
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  input_id text COLLATE pg_catalog."default" NOT NULL,
  value text COLLATE pg_catalog."default" NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
  OIDS = FALSE
);

-- index
CREATE UNIQUE INDEX IF NOT EXISTS input_id_idx
ON public.inputs USING btree
(input_id COLLATE pg_catalog."default")
TABLESPACE pg_default;

-- Trigger notification for messaging to PG Notify
CREATE FUNCTION notify_trigger() RETURNS trigger AS $trigger$
DECLARE
  rec RECORD;
  payload TEXT;
  column_name TEXT;
  column_value TEXT;
  payload_items TEXT[];
BEGIN
  -- Set record row depending on operation
  CASE TG_OP
  WHEN 'INSERT', 'UPDATE' THEN
    rec := NEW;
  WHEN 'DELETE' THEN
    rec := OLD;
  ELSE
    RAISE EXCEPTION 'Unknown TG_OP: "%". Should not occur!', TG_OP;
  END CASE;

  -- Get required fields
  FOREACH column_name IN ARRAY TG_ARGV LOOP
    EXECUTE format('SELECT $1.%I::TEXT', column_name)
    INTO column_value
    USING rec;
    payload_items := array_append(payload_items, '"' || replace(column_name, '"', '\"') || '":"' || replace(column_value, '"', '\"') || '"');
  END LOOP;

  -- Build the payload
  payload := ''
    || '{'
    || '"timestamp":"' || CURRENT_TIMESTAMP                    || '",'
    || '"operation":"' || TG_OP                                || '",'
    || '"schema":"'    || TG_TABLE_SCHEMA                      || '",'
    || '"table":"'     || TG_TABLE_NAME                        || '",'
    || '"data":{'      || array_to_string(payload_items, ',')  || '}'
    || '}';

  -- Notify the channel
  PERFORM pg_notify('db_notifications', payload);

  RETURN rec;
END;
$trigger$ LANGUAGE plpgsql;

-- trigger
CREATE TRIGGER inputs_notify AFTER INSERT OR UPDATE OR DELETE ON inputs
FOR EACH ROW EXECUTE PROCEDURE notify_trigger(
  'id',
  'input_id',
  'value'
);

-- INSERT
INSERT INTO inputs(input_id, value) VALUES ('foo', '')
