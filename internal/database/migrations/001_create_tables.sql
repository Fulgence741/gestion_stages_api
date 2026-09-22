--
-- PostgreSQL database dump
--

\restrict W2Vc1Geoz2kAZEMY7R7mqfNiYrOF9BREMynAyt2hCz9j5cXfaGugl8VSNixoBQV

-- Dumped from database version 14.24 (Ubuntu 14.24-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.24 (Ubuntu 14.24-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: etablissements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.etablissements (
    id_etablissement bigint NOT NULL,
    nom_etablissement character varying(150) NOT NULL,
    ville character varying(100) NOT NULL,
    adresse character varying(255)
);


ALTER TABLE public.etablissements OWNER TO postgres;

--
-- Name: etablissements_id_etablissement_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.etablissements ALTER COLUMN id_etablissement ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.etablissements_id_etablissement_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: etudiants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.etudiants (
    id_etudiant bigint NOT NULL,
    matricule character varying(50) NOT NULL,
    nom character varying(100) NOT NULL,
    prenom character varying(100) NOT NULL,
    telephone character varying(30),
    id_user bigint NOT NULL,
    id_etablissement bigint NOT NULL,
    id_filiere bigint NOT NULL
);


ALTER TABLE public.etudiants OWNER TO postgres;

--
-- Name: etudiants_id_etudiant_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.etudiants ALTER COLUMN id_etudiant ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.etudiants_id_etudiant_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: evaluations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.evaluations (
    id_evaluation bigint NOT NULL,
    note numeric(5,2) NOT NULL,
    commentaire text,
    date_evaluation timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id_rapport bigint NOT NULL,
    id_user bigint NOT NULL,
    CONSTRAINT evaluations_note_check CHECK (((note >= (0)::numeric) AND (note <= (20)::numeric)))
);


ALTER TABLE public.evaluations OWNER TO postgres;

--
-- Name: evaluations_id_evaluation_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.evaluations ALTER COLUMN id_evaluation ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.evaluations_id_evaluation_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: filieres; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.filieres (
    id_filiere bigint NOT NULL,
    nom_filiere character varying(150) NOT NULL
);


ALTER TABLE public.filieres OWNER TO postgres;

--
-- Name: filieres_id_filiere_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.filieres ALTER COLUMN id_filiere ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.filieres_id_filiere_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: maitre_stage; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.maitre_stage (
    id_maitre_stage bigint NOT NULL,
    nom character varying(100) NOT NULL,
    prenom character varying(100) NOT NULL,
    fonction character varying(100),
    telephone character varying(30),
    email character varying(255),
    id_structure bigint NOT NULL
);


ALTER TABLE public.maitre_stage OWNER TO postgres;

--
-- Name: maitre_stage_id_maitre_stage_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.maitre_stage ALTER COLUMN id_maitre_stage ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.maitre_stage_id_maitre_stage_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: proposer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.proposer (
    id_etablissement bigint NOT NULL,
    id_filiere bigint NOT NULL,
    date_proposition date DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.proposer OWNER TO postgres;

--
-- Name: rapports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rapports (
    id_rapport bigint NOT NULL,
    titre_rapport character varying(200) NOT NULL,
    chemin_fichier character varying(500) NOT NULL,
    date_depot timestamp with time zone,
    statut character varying(30) NOT NULL,
    note_finale numeric(5,2),
    id_stage bigint NOT NULL,
    CONSTRAINT rapports_note_finale_check CHECK (((note_finale IS NULL) OR ((note_finale >= (0)::numeric) AND (note_finale <= (20)::numeric))))
);


ALTER TABLE public.rapports OWNER TO postgres;

--
-- Name: rapports_id_rapport_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.rapports ALTER COLUMN id_rapport ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.rapports_id_rapport_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stages (
    id_stage bigint NOT NULL,
    intitule_stage character varying(200) NOT NULL,
    description text,
    date_debut date NOT NULL,
    date_fin date,
    niveau_etude character varying(50) NOT NULL,
    id_etudiant bigint NOT NULL,
    id_maitre_stage bigint NOT NULL,
    CONSTRAINT stages_check CHECK (((date_fin IS NULL) OR (date_fin >= date_debut)))
);


ALTER TABLE public.stages OWNER TO postgres;

--
-- Name: stages_id_stage_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.stages ALTER COLUMN id_stage ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.stages_id_stage_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: structures; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.structures (
    id_structure bigint NOT NULL,
    nom_structure character varying(150) NOT NULL,
    ville character varying(100) NOT NULL,
    adresse character varying(255),
    telephone character varying(30),
    email character varying(255)
);


ALTER TABLE public.structures OWNER TO postgres;

--
-- Name: structures_id_structure_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.structures ALTER COLUMN id_structure ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.structures_id_structure_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id_user bigint NOT NULL,
    nom character varying(100) NOT NULL,
    prenom character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    mot_de_passe character varying(255) NOT NULL,
    role character varying(30) NOT NULL,
    statut character varying(30) NOT NULL,
    date_creation timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id_user ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_user_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: etablissements etablissements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etablissements
    ADD CONSTRAINT etablissements_pkey PRIMARY KEY (id_etablissement);


--
-- Name: etudiants etudiants_id_user_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etudiants
    ADD CONSTRAINT etudiants_id_user_key UNIQUE (id_user);


--
-- Name: etudiants etudiants_matricule_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etudiants
    ADD CONSTRAINT etudiants_matricule_key UNIQUE (matricule);


--
-- Name: etudiants etudiants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etudiants
    ADD CONSTRAINT etudiants_pkey PRIMARY KEY (id_etudiant);


--
-- Name: evaluations evaluations_id_user_id_rapport_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluations
    ADD CONSTRAINT evaluations_id_user_id_rapport_key UNIQUE (id_user, id_rapport);


--
-- Name: evaluations evaluations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluations
    ADD CONSTRAINT evaluations_pkey PRIMARY KEY (id_evaluation);


--
-- Name: filieres filieres_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.filieres
    ADD CONSTRAINT filieres_pkey PRIMARY KEY (id_filiere);


--
-- Name: maitre_stage maitre_stage_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.maitre_stage
    ADD CONSTRAINT maitre_stage_pkey PRIMARY KEY (id_maitre_stage);


--
-- Name: proposer proposer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.proposer
    ADD CONSTRAINT proposer_pkey PRIMARY KEY (id_etablissement, id_filiere);


--
-- Name: rapports rapports_id_stage_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rapports
    ADD CONSTRAINT rapports_id_stage_key UNIQUE (id_stage);


--
-- Name: rapports rapports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rapports
    ADD CONSTRAINT rapports_pkey PRIMARY KEY (id_rapport);


--
-- Name: stages stages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stages
    ADD CONSTRAINT stages_pkey PRIMARY KEY (id_stage);


--
-- Name: structures structures_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.structures
    ADD CONSTRAINT structures_pkey PRIMARY KEY (id_structure);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id_user);


--
-- Name: etudiants etudiants_id_etablissement_id_filiere_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etudiants
    ADD CONSTRAINT etudiants_id_etablissement_id_filiere_fkey FOREIGN KEY (id_etablissement, id_filiere) REFERENCES public.proposer(id_etablissement, id_filiere) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: etudiants etudiants_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.etudiants
    ADD CONSTRAINT etudiants_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id_user) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: evaluations evaluations_id_rapport_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluations
    ADD CONSTRAINT evaluations_id_rapport_fkey FOREIGN KEY (id_rapport) REFERENCES public.rapports(id_rapport) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: evaluations evaluations_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluations
    ADD CONSTRAINT evaluations_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id_user) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: maitre_stage maitre_stage_id_structure_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.maitre_stage
    ADD CONSTRAINT maitre_stage_id_structure_fkey FOREIGN KEY (id_structure) REFERENCES public.structures(id_structure) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: proposer proposer_id_etablissement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.proposer
    ADD CONSTRAINT proposer_id_etablissement_fkey FOREIGN KEY (id_etablissement) REFERENCES public.etablissements(id_etablissement) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: proposer proposer_id_filiere_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.proposer
    ADD CONSTRAINT proposer_id_filiere_fkey FOREIGN KEY (id_filiere) REFERENCES public.filieres(id_filiere) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: rapports rapports_id_stage_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rapports
    ADD CONSTRAINT rapports_id_stage_fkey FOREIGN KEY (id_stage) REFERENCES public.stages(id_stage) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: stages stages_id_etudiant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stages
    ADD CONSTRAINT stages_id_etudiant_fkey FOREIGN KEY (id_etudiant) REFERENCES public.etudiants(id_etudiant) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: stages stages_id_maitre_stage_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stages
    ADD CONSTRAINT stages_id_maitre_stage_fkey FOREIGN KEY (id_maitre_stage) REFERENCES public.maitre_stage(id_maitre_stage) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

\unrestrict W2Vc1Geoz2kAZEMY7R7mqfNiYrOF9BREMynAyt2hCz9j5cXfaGugl8VSNixoBQV

