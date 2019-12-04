create table if not exists integracao_promedico
(
  cd_integracao      serial not null
    constraint integracao_promedico_pk
      primary key,
  accession_number   integer,
  cd_paciente        integer,
  ds_paciente        varchar(255),
  dt_nascimento      date,
  dt_exame           timestamp,
  ds_solicitante     varchar(255),
  ds_crm_solicitante varchar(255),
  cd_oficial         varchar(255),
  created_at         timestamp,
  ds_modalidade      varchar(2),
  ds_exame           varchar(255),
  ds_sexo            char,
  cd_medico          integer,
  cd_procedimento    varchar(255),
  sn_integrado       boolean default false,
  sn_cancelado       boolean default false
);

alter table integracao_promedico
  owner to dicomvix;

create unique index if not exists integracao_promedico_accession_number_uindex
  on integracao_promedico (accession_number);

