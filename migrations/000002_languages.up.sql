CREATE TABLE IF NOT EXISTS languages
(
    id             BIGSERIAL PRIMARY KEY,
    code           text NOT NULL,
    name           text NOT NULL,
    image_resource text NOT NULL,
    choose_title   text NOT NULL
);

INSERT INTO languages (code, name, image_resource, choose_title)
VALUES ('en', 'English', 'uk_flag', 'Please select the language you want to learn'),
       ('de', 'Deutsch', 'germany_flag', 'Bitte wählen Sie die Sprache aus, die Sie lernen möchten'),
       ('es', 'Español', 'spain_flag', 'Por favor, selecciona el lenguaje que deseas aprender'),
       ('ca', 'Català', 'catalonia_flag', 'Si us plau, selecciona el llenguatge que vols aprendre'),
       ('fr', 'Français', 'france_flag', 'Veuillez sélectionner la langue que vous apprendre');
