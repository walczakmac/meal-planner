-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table meal_category
(
    id   smallserial primary key,
    name varchar not null
);

insert into meal_category(name)
values ('sniadanie'),
    ('2-sniadanie'),
    ('obiad'),
    ('kolacja');

create table product_category
(
    id   smallserial primary key,
    name varchar not null
);

insert into product_category(name)
values ('owoce i warzywa'),
    ('nabiał'),
    ('zbożowe'),
    ('orzechy i ziarna'),
    ('inne'),
    ('mięso'),
    ('przyprawy i zioła'),
    ('pieczywo'),
    ('ryby i owoce morza'),
    ('napoje'),
    ('tłuszcze');

create table product
(
    id                  smallserial primary key,
    product_category_id smallint not null,
    name                varchar  not null
);

insert into product(product_category_id, name)
values
-- OWOCE I WARZYWA
    (1, 'Ananas'),
    (1, 'Awokado'),
    (1, 'Bakłażan'),
    (1, 'Banan'),
    (1, 'Bazylia (świeża)'),
    (1, 'Borówki amerykańskie'),
    (1, 'Brokuły'),
    (1, 'Brokuły, mrożone'),
    (1, 'Brzoskwinia'),
    (1, 'Buraczki wiórki'),
    (1, 'Burak'),
    (1, 'Ciecierzyca (w zalewie)'),
    (1, 'Cukinia'),
    (1, 'Cytryna'),
    (1, 'Ćwikła z chrzanem'),
    (1, 'Daktyle, suszone'),
    (1, 'Fasola czerwona w zalewie (konserwowa)'),
    (1, 'Fasolka szparagowa, mrożona'),
    (1, 'Groszek zielony, mrożony'),
    (1, 'Gruszka'),
    (1, 'Jabłko'),
    (1, 'Kalafior, mrożony'),
    (1, 'Kapusta, kiszona'),
    (1, 'Kiwi'),
    (1, 'Koper ogrodowy'),
    (1, 'Kukurydza, konserwowa'),
    (1, 'Limonka'),
    (1, 'Maliny'),
    (1, 'Maliny, mrożone'),
    (1, 'Mandarynki'),
    (1, 'Mango'),
    (1, 'Marchew'),
    (1, 'MIX sałat'),
    (1, 'Morele'),
    (1, 'Morele, suszone'),
    (1, 'Nektarynka'),
    (1, 'Ogórek'),
    (1, 'Ogórek zielony (długi)'),
    (1, 'Ogórki, kiszone'),
    (1, 'Oliwki czarne'),
    (1, 'Oliwki zielone, marynowane, konserwowe'),
    (1, 'Papryczka ostra (chili)'),
    (1, 'Papryka czerwona'),
    (1, 'Pieczarka uprawna, świeża'),
    (1, 'Pietruszka, korzeń'),
    (1, 'Pietruszka, liście'),
    (1, 'Pomarańcza'),
    (1, 'Pomidor'),
    (1, 'Pomidory koktajlowe'),
    (1, 'Pomidory z puszki'),
    (1, 'Pomidory z puszki (krojone)'),
    (1, 'Por'),
    (1, 'Przecier pomidorowy'),
    (1, 'Rukola'),
    (1, 'Rzodkiewka'),
    (1, 'Sałata lodowa'),
    (1, 'Soczewica czerwona, nasiona suche'),
    (1, 'Suszone pomidory'),
    (1, 'Szpinak'),
    (1, 'Szpinak, mrożony'),
    (1, 'Truskawki'),
    (1, 'Truskawki, mrożone'),
    (1, 'Warzywa na patelnię chińskie'),
    (1, 'Warzywne trio mrożone: brokuł, marchew, kalafior'),
    (1, 'Ziemniaki'),
    (1, 'Żurawina suszona'),
    (1, 'Czosnek'),
-- NABIAŁ
    (2, 'Jaja kurze całe'),
    (2, 'Jogurt naturalny 2%'),
    (2, 'Kefir'),
    (2, 'Mleko spożywcze 2%'),
    (2, 'Mozzarella light'),
    (2, 'Ser twarogowy chudy'),
    (2, 'Ser feta - Favita 12%'),
    (2, 'Ser, gouda tłusty'),
    (2, 'Serek wiejski light'),
    (2, 'Skyr naturalny'),
    (2, 'Skyr owocowy'),
    (2, 'Śmietana 12%'),
    (2, 'Śmietana 18%'),
-- ZBOŻOWE
    (3, 'Kasza bulgur'),
    (3, 'Kasza gryczana'),
    (3, 'Kasza jaglana'),
    (3, 'Kasza jęczmienna, pęczak'),
    (3, 'Kasza manna'),
    (3, 'Komosa ryżowa'),
    (3, 'Makaron pełnoziarnisty'),
    (3, 'Makaron penne (pełnoziarnisty)'),
    (3, 'Makaron ryżowy'),
    (3, 'Makaron spaghetti pełnoziarnisty'),
    (3, 'Mąka pszenna (typ 2000, pełnoziarnista)'),
    (3, 'Płatki jaglane'),
    (3, 'Płatki owsiane'),
    (3, 'Ryż basmati'),
    (3, 'Ryż brązowy'),
    (3, 'Skrobia ziemniaczana'),
    (3, 'Wafle ryżowe naturalne'),
-- ORZECHY I ZIARNA
    (4, 'Dynia, pestki, łuskane'),
    (4, 'Kakao'),
    (4, 'Migdały'),
    (4, 'Migdały w płatkach'),
    (4, 'Nasiona chia'),
    (4, 'Orzechy brazylijskie'),
    (4, 'Orzechy nerkowca (bez soli)'),
    (4, 'Orzechy włoskie'),
    (4, 'Słonecznik, nasiona, łuskane'),
    (4, 'Tahini'),
    (4, 'Wiórki kokosowe'),
-- INNE
    (5, 'Bulion warzywny (domowy)'),
    (5, 'Czekolada gorzka'),
    (5, 'Dżem 100% owoców'),
    (5, 'Hummus klasyczny'),
    (5, 'Keczup'),
    (5, 'Kostka ekologiczny bulion drobiowy'),
    (5, 'Przyprawa korzenna do piernika (bez cukru)'),
    (5, 'Miód pszczeli'),
    (5, 'Musztarda'),
    (5, 'Ocet jabłkowy'),
    (5, 'Odżywka białkowa'),
    (5, 'Płatki drożdżowe'),
    (5, 'Proszek do pieczenia'),
    (5, 'Słodzik'),
    (5, 'Sos sojowy ciemny'),
    (5, 'Tofu naturalne'),
    (5, 'Tofu wędzone'),
-- MIĘSO
    (6, 'Chude mięso mielone z szynki wieprzowej'),
    (6, 'Filet z kurczaka (wędzony)'),
    (6, 'Mielony filet z piersi indyka (bez skóry)'),
    (6, 'Mielony filet z piersi kurczaka (bez skóry)'),
    (6, 'Mięso mielone z piersi kurczaka'),
    (6, 'Mięso z piersi indyka, bez skóry'),
    (6, 'Mięso z piersi kurczaka, bez skóry'),
    (6, 'Mięso z ud kurczaka, bez skóry'),
    (6, 'Polędwica wieprzowa (surowa)'),
    (6, 'Szynka z piersi kurczaka'),
-- PRZYPRAWY I ZIOŁA
    (7, 'Bazylia (suszona)'),
    (7, 'Cynamon'),
    (7, 'Kmin rzymski (kumin)'),
    (7, 'Kostka rosołowa warzywna'),
    (7, 'Kostka rosołowa warzywna'),
    (7, 'Kurkuma'),
    (7, 'Oregano (suszone)'),
    (7, 'Tymianek'),
-- PIECZYWO
    (8, 'Bułki grahamki'),
    (8, 'Bułki pszenne zwykłe'),
    (8, 'Chleb żytni razowy'),
    (8, 'Tortilla pszenna (duża)'),
-- RYBY I OWOCE MORZA
    (9, 'Dorsz, świeży'),
    (9, 'Łosoś, wędzony'),
    (9, 'Łosoś, świeży'),
    (9, 'Makrela, wędzona'),
    (9, 'Pstrąg tęczowy, świeży'),
    (9, 'Tuńczyk w sosie własnym'),
-- NAPOJE
    (10, 'Sok marchwiowy'),
    (10, 'Sok cytrynowy'),
-- TŁUSZCZE
    (11, 'Oliwa z oliwek'),
    (11, 'Olej sezamowy');

create table ingredient
(
    id              smallserial primary key,
    meal_variant_id smallint not null,
    product_id      smallint not null,
    amount          smallint not null,
    unit            varchar  not null,
    snack           boolean  not null default false
);

create table macro
(
    id              smallserial primary key,
    meal_variant_id smallint not null,
    proteins        real     not null,
    fats            real     not null,
    carbs           real     not null,
    fiber           real     not null
);

create table meal
(
    id               smallserial primary key,
    meal_category_id smallint not null,
    "name"           varchar  not null,
    description      varchar  not null,
    "day"            smallint not null
);

create table meal_variant
(
    id         smallserial primary key,
    meal_id    smallint not null,
    kcal       real     not null,
    kcal_daily smallint not null,
    person     varchar  not null
);

create table plan
(
    id smallserial primary key,
    date date not null
);

create table plan_meal
(
    id smallserial primary key,
    plan_id smallint not null,
    meal_variant_id smallint not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop table if exists ingredient;
drop table if exists product;
drop table if exists macro;
drop table if exists meal;
drop table if exists meal_category;
drop table if exists meal_variant;
drop table if exists product_category;
drop table if exists plan_meal;
drop table if exists plan;

-- +goose StatementEnd
