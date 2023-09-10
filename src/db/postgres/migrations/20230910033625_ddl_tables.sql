-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS hk
    AUTHORIZATION postgres;
CREATE TABLE "hk"."access_rol" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar
);

CREATE TABLE "hk"."user_access_rol" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "access_rol_id" uuid,
  "users_id" uuid
);

CREATE TABLE "hk"."type_user" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar
);

CREATE TABLE "hk"."users_status" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar
);

CREATE TABLE "hk"."credentials" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "password" varchar,
  "email" varchar,
  "status_id" uuid
);

CREATE TABLE "hk"."week_day" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_schedule_id" uuid,
  "name" varchar,
  "day" varchar,
  "month" varchar,
  "year" varchar,
  "start_time" time,
  "end_time" time
);

CREATE TABLE "hk"."business_schedule" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_locations_id" uuid
);

CREATE TABLE "hk"."offers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "sales_off" uuid,
  "discount_level" decimal(5,2)
);

CREATE TABLE "hk"."services_offers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "offers_id" uuid,
  "services_id" uuid
);

CREATE TABLE "hk"."sales_off" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_locations_id" uuid,
  "title" varchar,
  "start_time" timestamp,
  "end_time" timestamp,
  "active" bool
);

CREATE TABLE "hk"."prices" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "price" decimal(10,2),
  "services_id" uuid
);

CREATE TABLE "hk"."services" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_locations_id" uuid,
  "name" varchar,
  "time" time
);

CREATE TABLE "hk"."employees" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_locations_id" uuid,
  "employee_id" uuid
);

CREATE TABLE "hk"."business_locations" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar,
  "address" varchar,
  "description" varchar,
  "barrio_id" uuid,
  "business_id" uuid
);

CREATE TABLE "hk"."country" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."department" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "country_id" uuid,
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."municipio" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "department_id" uuid,
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."corregimientos" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "municipio_id" uuid,
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."localidades" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "municipio_id" uuid,
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."barrios" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "municipio_id" uuid,
  "localidad_id" uuid,
  "code" varchar,
  "name" varchar
);

CREATE TABLE "hk"."business" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner_id" uuid,
  "name" varchar,
  "description" varchar
);

CREATE TABLE "hk"."type_multimedia" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "type" varchar
);

CREATE TABLE "hk"."multimedia" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "link" varchar,
  "assets_id" uuid,
  "type" uuid
);

CREATE TABLE "hk"."assets" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "business_locations_id" uuid,
  "sales_off_id" uuid
);

CREATE TABLE "hk"."type_social_media" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar,
  "link_base" varchar
);

CREATE TABLE "hk"."social_media" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "type_social_media_id" uuid,
  "link" varchar,
  "business_locations_id" uuid
);

CREATE TABLE "hk"."calification" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "users_id" uuid,
  "reservation_id" uuid,
  "rate" smallint,
  "note" varchar,
  "date_rate" timestamp DEFAULT (now())
);

CREATE TABLE "hk"."reservation" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "code" smallint,
  "note" varchar,
  "start_time" timestamp,
  "end_time" timestamp,
  "services_id" uuid,
  "business_locations_id" uuid
);

CREATE TABLE "hk"."client_reservation" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "client_id" uuid,
  "reservation_id" uuid
);

CREATE TABLE "hk"."reservation_assigned_employee" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "employee_id" uuid,
  "reservation_id" uuid
);

CREATE TABLE "hk"."type_status_reservation" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar
);

CREATE TABLE "hk"."status_reservation" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "type_status_reservation_id" uuid,
  "reservation_id" uuid
);

CREATE TABLE "hk"."history_reservation" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "status_reservation_id" uuid,
  "type_status_reservation_id" uuid,
  "create_at" timestamp DEFAULT (now())
);

CREATE TABLE "hk"."users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar,
  "credentials_id" uuid,
  "type_user_id" uuid
);

COMMENT ON COLUMN "hk"."access_rol"."name" IS 'nombre del rol';

COMMENT ON COLUMN "hk"."type_user"."name" IS 'tipo de usuario: 
                                cliente: usuario consumidor (tipo inicial)
                                empleado: usuario prestador de servicio
                                owner: administrador/dueño de local';

COMMENT ON COLUMN "hk"."users_status"."name" IS 'estados del acceso al usuario: 
                            active: usuario activo
                            pending: usuario en espera de validar correo
                            block: usuario bloqueado
                            delete: usuario dado de baja';

COMMENT ON COLUMN "hk"."credentials"."password" IS 'password usuario';

COMMENT ON COLUMN "hk"."credentials"."email" IS 'correo usuario';

COMMENT ON COLUMN "hk"."credentials"."status_id" IS 'referencia a tabla users_status';

COMMENT ON COLUMN "hk"."week_day"."name" IS 'nombre de la agenda ej: jornada normal, extendida, feriados o fin de semana';

COMMENT ON COLUMN "hk"."week_day"."day" IS 'dia';

COMMENT ON COLUMN "hk"."week_day"."month" IS 'mes';

COMMENT ON COLUMN "hk"."week_day"."year" IS 'año';

COMMENT ON COLUMN "hk"."week_day"."start_time" IS 'hora de inicio';

COMMENT ON COLUMN "hk"."week_day"."end_time" IS 'hora de fin';

COMMENT ON COLUMN "hk"."offers"."discount_level" IS 'nivel de descuento ej: 50, 12 o 30 por ciento';

COMMENT ON COLUMN "hk"."sales_off"."title" IS 'nombre de la promoción ej: ofertas de fin de año';

COMMENT ON COLUMN "hk"."sales_off"."start_time" IS 'fecha de inicio';

COMMENT ON COLUMN "hk"."sales_off"."end_time" IS 'fecha de inicio';

COMMENT ON COLUMN "hk"."sales_off"."active" IS 'indicador de estado';

COMMENT ON COLUMN "hk"."prices"."price" IS 'precio del servicio';

COMMENT ON COLUMN "hk"."services"."name" IS 'nombre del servicio';

COMMENT ON COLUMN "hk"."services"."time" IS 'hora aproximada que se demora el servicio';

COMMENT ON COLUMN "hk"."employees"."business_locations_id" IS 'empresa empleadora';

COMMENT ON COLUMN "hk"."employees"."employee_id" IS 'id de la tabla usuario';

COMMENT ON COLUMN "hk"."business_locations"."name" IS 'nombre de la sucursal';

COMMENT ON COLUMN "hk"."business_locations"."address" IS 'dirección del local';

COMMENT ON COLUMN "hk"."business_locations"."description" IS 'descripción adicional';

COMMENT ON COLUMN "hk"."business_locations"."business_id" IS 'empresa padre';

COMMENT ON COLUMN "hk"."country"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."country"."name" IS 'nombre del país';

COMMENT ON COLUMN "hk"."department"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."department"."name" IS 'nombre del departamento';

COMMENT ON COLUMN "hk"."municipio"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."municipio"."name" IS 'nombre del municipio/ciudad';

COMMENT ON COLUMN "hk"."corregimientos"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."corregimientos"."name" IS 'nombre del corregimientos';

COMMENT ON COLUMN "hk"."localidades"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."localidades"."name" IS 'nombre del localidad';

COMMENT ON COLUMN "hk"."barrios"."code" IS 'código dane';

COMMENT ON COLUMN "hk"."barrios"."name" IS 'nombre del barrio';

COMMENT ON COLUMN "hk"."business"."name" IS 'nombre de la empresa principal';

COMMENT ON COLUMN "hk"."business"."description" IS 'descripción adicional';

COMMENT ON COLUMN "hk"."type_multimedia"."type" IS 'tipo de archivo';

COMMENT ON COLUMN "hk"."multimedia"."link" IS 'enlace de acceso al archivo';

COMMENT ON COLUMN "hk"."assets"."business_locations_id" IS 'referencia de la empresa';

COMMENT ON COLUMN "hk"."assets"."sales_off_id" IS 'referencia de la oferta si aplica';

COMMENT ON COLUMN "hk"."type_social_media"."name" IS 'red social ej: facebook';

COMMENT ON COLUMN "hk"."type_social_media"."link_base" IS 'link base de red social';

COMMENT ON COLUMN "hk"."social_media"."type_social_media_id" IS 'referencia a la red social especifica';

COMMENT ON COLUMN "hk"."social_media"."link" IS 'link perfil red social';

COMMENT ON COLUMN "hk"."social_media"."business_locations_id" IS 'referencia de la empresa';

COMMENT ON COLUMN "hk"."calification"."users_id" IS 'referencia del cliente';

COMMENT ON COLUMN "hk"."calification"."reservation_id" IS 'referencia de la reserva';

COMMENT ON COLUMN "hk"."calification"."rate" IS 'rango de la calificación de 1 a 5';

COMMENT ON COLUMN "hk"."calification"."note" IS 'nota o comentario';

COMMENT ON COLUMN "hk"."calification"."date_rate" IS 'fecha y hora de la calificación';

COMMENT ON COLUMN "hk"."reservation"."code" IS 'código de la reserva de 4 dígitos';

COMMENT ON COLUMN "hk"."reservation"."note" IS 'nota o comentario de la reserva';

COMMENT ON COLUMN "hk"."reservation"."start_time" IS 'inicio: hora y fecha del servicio';

COMMENT ON COLUMN "hk"."reservation"."end_time" IS 'fin: hora y fecha del servicio';

COMMENT ON COLUMN "hk"."reservation"."services_id" IS 'referencia del servicio';

COMMENT ON COLUMN "hk"."reservation"."business_locations_id" IS 'referencia de la empresa';

ALTER TABLE "hk"."user_access_rol" ADD FOREIGN KEY ("access_rol_id") REFERENCES "hk"."access_rol" ("id");

ALTER TABLE "hk"."user_access_rol" ADD FOREIGN KEY ("users_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."credentials" ADD FOREIGN KEY ("status_id") REFERENCES "hk"."users_status" ("id");

ALTER TABLE "hk"."week_day" ADD FOREIGN KEY ("business_schedule_id") REFERENCES "hk"."business_schedule" ("id");

ALTER TABLE "hk"."business_schedule" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."offers" ADD FOREIGN KEY ("sales_off") REFERENCES "hk"."sales_off" ("id");

ALTER TABLE "hk"."services_offers" ADD FOREIGN KEY ("offers_id") REFERENCES "hk"."offers" ("id");

ALTER TABLE "hk"."services_offers" ADD FOREIGN KEY ("services_id") REFERENCES "hk"."services" ("id");

ALTER TABLE "hk"."sales_off" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."prices" ADD FOREIGN KEY ("services_id") REFERENCES "hk"."services" ("id");

ALTER TABLE "hk"."services" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."employees" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."employees" ADD FOREIGN KEY ("employee_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."business_locations" ADD FOREIGN KEY ("barrio_id") REFERENCES "hk"."barrios" ("id");

ALTER TABLE "hk"."business_locations" ADD FOREIGN KEY ("business_id") REFERENCES "hk"."business" ("id");

ALTER TABLE "hk"."department" ADD FOREIGN KEY ("country_id") REFERENCES "hk"."country" ("id");

ALTER TABLE "hk"."municipio" ADD FOREIGN KEY ("department_id") REFERENCES "hk"."department" ("id");

ALTER TABLE "hk"."corregimientos" ADD FOREIGN KEY ("municipio_id") REFERENCES "hk"."municipio" ("id");

ALTER TABLE "hk"."localidades" ADD FOREIGN KEY ("municipio_id") REFERENCES "hk"."municipio" ("id");

ALTER TABLE "hk"."barrios" ADD FOREIGN KEY ("municipio_id") REFERENCES "hk"."municipio" ("id");

ALTER TABLE "hk"."barrios" ADD FOREIGN KEY ("localidad_id") REFERENCES "hk"."localidades" ("id");

ALTER TABLE "hk"."business" ADD FOREIGN KEY ("owner_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."multimedia" ADD FOREIGN KEY ("assets_id") REFERENCES "hk"."assets" ("id");

ALTER TABLE "hk"."multimedia" ADD FOREIGN KEY ("type") REFERENCES "hk"."type_multimedia" ("id");

ALTER TABLE "hk"."assets" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."assets" ADD FOREIGN KEY ("sales_off_id") REFERENCES "hk"."sales_off" ("id");

ALTER TABLE "hk"."social_media" ADD FOREIGN KEY ("type_social_media_id") REFERENCES "hk"."type_social_media" ("id");

ALTER TABLE "hk"."social_media" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."calification" ADD FOREIGN KEY ("users_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."calification" ADD FOREIGN KEY ("reservation_id") REFERENCES "hk"."reservation" ("id");

ALTER TABLE "hk"."reservation" ADD FOREIGN KEY ("services_id") REFERENCES "hk"."services" ("id");

ALTER TABLE "hk"."reservation" ADD FOREIGN KEY ("business_locations_id") REFERENCES "hk"."business_locations" ("id");

ALTER TABLE "hk"."client_reservation" ADD FOREIGN KEY ("client_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."client_reservation" ADD FOREIGN KEY ("reservation_id") REFERENCES "hk"."reservation" ("id");

ALTER TABLE "hk"."reservation_assigned_employee" ADD FOREIGN KEY ("employee_id") REFERENCES "hk"."users" ("id");

ALTER TABLE "hk"."reservation_assigned_employee" ADD FOREIGN KEY ("reservation_id") REFERENCES "hk"."reservation" ("id");

ALTER TABLE "hk"."status_reservation" ADD FOREIGN KEY ("type_status_reservation_id") REFERENCES "hk"."type_status_reservation" ("id");

ALTER TABLE "hk"."status_reservation" ADD FOREIGN KEY ("reservation_id") REFERENCES "hk"."reservation" ("id");

ALTER TABLE "hk"."history_reservation" ADD FOREIGN KEY ("status_reservation_id") REFERENCES "hk"."status_reservation" ("id");

ALTER TABLE "hk"."history_reservation" ADD FOREIGN KEY ("type_status_reservation_id") REFERENCES "hk"."type_status_reservation" ("id");

ALTER TABLE "hk"."users" ADD FOREIGN KEY ("credentials_id") REFERENCES "hk"."credentials" ("id");

ALTER TABLE "hk"."users" ADD FOREIGN KEY ("type_user_id") REFERENCES "hk"."type_user" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS hk CASCADE;
-- +goose StatementEnd
