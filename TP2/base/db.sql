-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: pbbv
-- ------------------------------------------------------
-- Server version	8.0.37

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cursos`
--

DROP TABLE IF EXISTS `cursos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cursos` (
                          `id_curso` int NOT NULL AUTO_INCREMENT,
                          `Titulo` varchar(100) DEFAULT NULL,
                          `Fecha_inicio` date NOT NULL,
                          `Categoria` varchar(25) NOT NULL,
                          `Archivo` varchar(50) DEFAULT NULL,
                          `Descripcion` text NOT NULL,
                          PRIMARY KEY (`id_curso`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cursos`
--

LOCK TABLES `cursos` WRITE;
/*!40000 ALTER TABLE `cursos` DISABLE KEYS */;
INSERT INTO `cursos` VALUES
                         (1,'IntroducciÃ³n a la ProgramaciÃ³n','2024-06-01','InformÃ¡tica','introduccion_programacion.pdf','Curso introductorio sobre programaciÃ³n para principiantes.'),
                         (3,'InglÃ©s Conversacional','2024-07-01','Idiomas','ingles_conversacional.mp4','Curso diseÃ±ado para mejorar las habilidades de conversaciÃ³n en inglÃ©s.'),
                         (4,'DiseÃ±o GrÃ¡fico Profesional','2024-07-15','DiseÃ±o','diseno_grafico_profesional.zip','Curso completo sobre diseÃ±o grÃ¡fico utilizando herramientas profesionales.'),
                         (5,'Finanzas Personales','2024-08-01','Finanzas',NULL,'Curso prÃ¡ctico para gestionar de manera efectiva las finanzas personales.'),
                         (6,'Cocina MediterrÃ¡nea','2024-08-15','Cocina','cocina_mediterranea_recetas.docx','Curso de cocina centrado en recetas saludables de la cocina mediterrÃ¡nea.'),
                         (7,'Desarrollo Web con HTML y CSS','2024-09-01','InformÃ¡tica','desarrollo_web_html_css.pdf','Curso introductorio sobre desarrollo web utilizando HTML y CSS.'),
                         (8,'GestiÃ³n de Proyectos con Scrum','2024-09-15','Negocios','gestion_proyectos_scrum.pptx','Curso prÃ¡ctico sobre metodologÃ­as Ã¡giles de gestiÃ³n de proyectos.'),
                         (9,'FotografÃ­a Digital BÃ¡sica','2024-10-01','Arte','fotografia_digital_basica.jpg','Curso para principiantes que cubre los fundamentos de la fotografÃ­a digital.'),
                         (10,'IntroducciÃ³n a la Inteligencia Artificial','2024-10-15','TecnologÃ­a',NULL,'Curso introductorio que explora conceptos bÃ¡sicos de inteligencia artificial y sus aplicaciones.'),
                         (11,'Marketing en Redes Sociales','2024-11-01','Marketing','marketing_redes_sociales.mp4','Curso avanzado sobre estrategias de marketing en diversas plataformas de redes sociales.'),
                         (12,'Dibujo y Pintura para Principiantes','2024-11-15','Arte','dibujo_pintura_principiantes.docx','Curso diseÃ±ado para aquellos que desean aprender tÃ©cnicas bÃ¡sicas de dibujo y pintura.'),
                         (13,'ProgramaciÃ³n en Python','2024-12-01','InformÃ¡tica','programacion_python.zip','Curso completo sobre programaciÃ³n en Python, desde conceptos bÃ¡sicos hasta tÃ©cnicas avanzadas.'),
                         (14,'GestiÃ³n del Tiempo y Productividad','2024-12-15','Desarrollo Personal',NULL,'Curso prÃ¡ctico para mejorar la gestiÃ³n del tiempo y aumentar la productividad personal y profesional.'),
                         (15,'NutriciÃ³n y DietÃ©tica','2025-01-01','Salud','nutricion_dietetica.pdf','Curso que aborda los principios bÃ¡sicos de la nutriciÃ³n y ofrece pautas para mantener una dieta equilibrada.'),
                         (16,'Aprendizaje AutomÃ¡tico con TensorFlow','2025-01-15','TecnologÃ­a','aprendizaje_automatico_tensorflow.zip','Curso avanzado sobre aprendizaje automÃ¡tico utilizando la biblioteca TensorFlow de Google.'),
                         (19,'hola','2024-06-01','Nueva categorÃ­a','nuevo_archivo.pdf','Nueva descripciÃ³n del curso');
/*!40000 ALTER TABLE `cursos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inscripciones`
--

DROP TABLE IF EXISTS `inscripciones`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inscripciones` (
                                 `Id_usuario` int DEFAULT NULL,
                                 `Id_curso` int NOT NULL,
                                 `Fecha_inicio` date NOT NULL,
                                 `Comentario` text,
                                 KEY `fk_Curso` (`Id_curso`),
                                 KEY `fk_Users` (`Id_usuario`),
                                 CONSTRAINT `fk_Curso` FOREIGN KEY (`Id_curso`) REFERENCES `cursos` (`id_curso`),
                                 CONSTRAINT `fk_Users` FOREIGN KEY (`Id_usuario`) REFERENCES `users` (`Id_usuario`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inscripciones`
--

LOCK TABLES `inscripciones` WRITE;
/*!40000 ALTER TABLE `inscripciones` DISABLE KEYS */;
INSERT INTO `inscripciones` VALUES
                                (3,3,'2024-07-01','Quiero mejorar mi habilidad para hablar inglÃ©s en conversaciones.'),
                                (4,4,'2024-07-15','Me apasiona el diseÃ±o grÃ¡fico y quiero perfeccionar mis habilidades.'),
                                (5,5,'2024-08-01','Necesito mejorar mi gestiÃ³n financiera personal.'),
                                (6,6,'2024-08-15','Me encanta cocinar y quiero aprender mÃ¡s sobre la cocina mediterrÃ¡nea.'),
                                (7,7,'2024-09-01','Quiero empezar a desarrollar mis propias pÃ¡ginas web.'),
                                (8,8,'2024-09-15','Necesito mejorar la gestiÃ³n de proyectos en mi empresa.'),
                                (9,9,'2024-10-01','Siempre he querido aprender mÃ¡s sobre fotografÃ­a digital.'),
                                (10,10,'2024-10-15','Estoy intrigado por el potencial de la inteligencia artificial.'),
                                (11,11,'2024-11-01','Quiero mejorar mi estrategia de marketing en redes sociales para mi negocio.'),
                                (12,12,'2024-11-15','El dibujo y la pintura siempre me han interesado.'),
                                (13,13,'2024-12-01','Quiero aprender programaciÃ³n en Python para expandir mis habilidades.'),
                                (14,14,'2024-12-15','Necesito mejorar mi gestiÃ³n del tiempo y ser mÃ¡s productivo.'),
                                (15,15,'2025-01-01','Estoy interesado en aprender mÃ¡s sobre nutriciÃ³n y dietÃ©tica.'),
                                (16,16,'2025-01-15','Me gustarÃ­a adentrarme en el aprendizaje automÃ¡tico con TensorFlow.');
/*!40000 ALTER TABLE `inscripciones` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `Id_usuario` int NOT NULL AUTO_INCREMENT,
                         `Nombre_Usuario` varchar(25) NOT NULL,
                         `Nombre` varchar(25) NOT NULL,
                         `Apellido` varchar(25) NOT NULL,
                         `Email` varchar(50) NOT NULL,
                         `ContraseÃ±a` varchar(100) DEFAULT NULL,
                         `Tipo` varchar(20) DEFAULT NULL,
                         PRIMARY KEY (`Id_usuario`),
                         CONSTRAINT `users_chk_1` CHECK ((`Tipo` in (_utf8mb3'estudiante',_utf8mb3'administrador')))
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'pauliiortizz','Paulina','Ortiz Noseda','paulinaortizilr5@gmail.com','$2a$10$5FGmj62z24Uox.JqbKcERuld.PC/n.vjYrfEbj8gk/5v5riC.k1om','administrador'),
                           (2,'baujuncos','Bautista','Juncos','baujuncos1205@gmail.com','$2a$10$Z3dReNWUIU9wgLyTwVuBee1lTJk.qI7alX9CsbEKSev/N4Rgl7jFG','estudiante'),
                           (3,'belutreachii','Belen','Treachi','belentreachi@gmail.com','$2a$10$lDObM./7A5uUqPTFsxiZRe.eb/F9XcazQltlb5JXmWz3opR6B9UaC','estudiante'),
                           (4,'virrodriguez','Virginia','Rodriguez','virchu23@gmail.com','$2a$10$q7eF8ZvXR.ON8p.WOweezunQf3dAHhg0ioBpg1UpmGOiNTiuGU5zu','estudiante'),
                           (5,'magdanoseda','Magdalena','Noseda','maguinoseda@hotmail.com','$2a$10$N.fJr2BaMl61CaFOQzGN1OFX5j.O3VlX9DT41UeQpVtC/ows3l.Qi','administrador'),
                           (6,'olivetosofi','Sofia','Oliveto','oliveto078@gmail.com','$2a$10$XWVlWCtCDPw0LHjw0H7NEePvwbHWhJsVq76OqOjel7/XA9cWucszq','estudiante'),
                           (7,'juanperez','Juan','Perez','juanperez@mail.com','$2a$10$o/KrXhQ5Yd1xeTRpS2.uo.djq5tPlXCV2UdyeDYqIikMWyE6x.SY2','estudiante'),
                           (8,'anitaagomeeez','Ana','Gomez','anagomez@mail.com','$2a$10$VpeXje1aQcIk4Luvaw62HeD5k1m1gCw6GMrt4Z38FK1DHBCK6zH1W','Estudiante'),
                           (9,'carlitos4rodri','Carlos','Rodriguez','carlosrodriguez@mail.com','$2a$10$zdRD.yJQ0xc5HBy7UUd2rOwczH4klga/WlwOMlfngO31oCXtG2OT.','administrador'),
                           (10,'luchimartinez','Lucia','Martinez','luciamartinez@mail.com','$2a$10$cTlHN2LiX.HxLtYZEnPUuucGO8Hb/eRnjDrNZHvJzrDzI5VoHAoRy','estudiante'),
                           (11,'mariitaa','Maria','Fernandez','mariafernandez@mail.com','$2a$10$un7AGBQxFzEHFpkwZhb9suviMmZE0IocOs40ACtexUn6JBsl.xOG.','estudiante'),
                           (12,'peposanchez','Pedro','SÃ¡nchez','pedrosanchez@example.com','$2a$10$nsS51Rjhd3w8/f0oQf3CbeOp4TuTwx.D8yepTCq5dixVgtCsa3ymq','administrador'),
                           (13,'laurita','Laura','RodrÃ­guez','laurarodriguez@example.com','$2a$10$eBS8ZJkY05lwOC2C93nC.Ok1G.M4bK4GJRGRfqmM2VVbVmK.H5rHi','estudiante'),
                           (14,'fernandezsooo','SofÃ­a','FernÃ¡ndez','sofiafernandez@example.com','$2a$10$tCwh2xIRvmStJ/V0dhhxi.9.Gy9GhdeGf34a7EE68d48Ph9KPm8aK','estudiante'),
                           (15,'gomezmiguelito','Miguel','GÃ³mez','miguelgomez@example.com','$2a$10$6XV.mB5PaqoKewF3n6kjZejzoZK.NSqvrfh.Bu.q1ZvzYm7UPJs5K','estudiante'),
                           (16,'diazelena','Elena','DÃ­az','elenadiaz@example.com','$2a$10$9VPZDPKpUyGXvx7gWoi3GujKHNS80zfZY5PbJWlEg6P1oHzpTzCba','administrador'),
                           (17,'javi34alvarez','Javier','Alvarez','javieralvarez@example.com','$2a$10$Z7i8px9MnayE7xltNjDlh.IVrZaopocHw2qfHjnPgBl6d32PydN.2','estudiante'),
                           (18,'pauchii','Paula','Torres','paulatorres@example.com','$2a$10$dt1AROMAp6lcP3xchzptG.tFf35gcayipYV9bhFF41MHWOEsXw5me','estudiante'),
                           (19,'santoscarmelita','Carmen','Santos','carmensantos@example.com','$2a$10$MIsB4uBIhOqrepRqUbvfjewBUVmCTDD8GOR3/b2bsKpNIAYJowoe6','estudiante'),
                           (20,'alejimenez','Alejandro','JimÃ©nez','alejandrojimenez@example.com','$2a$10$mosE.G46MvGI3Y06jnuFxOC/xZuWyR3/k2Tq2LH2YdMwtKz97fflW','estudiante'),
                           (21,'morenorosi','Rosa','Moreno','rosamoreno@example.com','$2a$10$w4oZs06Ebe3T3CJnnxiUTu7tkH0r34xHol6t6KrCXCENPI7Cn/xiW','administrador'),
                           (22,'hernandez44','Pablo','HernÃ¡ndez','pablohernandez@example.com','$2a$10$yx0hbUWP8b6.m/xUzTcytu7YkEykuQUwotk5K.5U02pwP9Dvbqo2u','estudiante'),
                           (23,'lugarcia','Luisa','GarcÃ­a','luisagarcia@example.com','$2a$10$flJUQ2IfyRKJpqLvJl37Seqx.X6Q8cM0KwpLczPsbPtIkH9ugqVyq','estudiante'),
                           (24,'vargasfer','Fernando','Vargas','fernandovargas@example.com','$2a$10$HvDvSTnYb/8tOu9EItjtl./YwojR4ByX1fM00DDR/xRZ6km3DZwru','estudiante'),
                           (25,'luchi_lopez','LucÃ­a','LÃ³pez','lucialopez@example.com','$2a$10$bVTB56xZ6usE1siCYwp.kunhpJ94uQX9Vh1gW8Z2WC51MkUQvVU2q','administrador'),
                           (26,'pachii__','Paz','Martinez','pachita@gmail.com','34b85dffdbd56419d1d8c666648be993770718bc','estudiante'),
                           (27,'Marimarparma','Marta','Parma','martaparma@gmail.com','a846e8344cbd3ff5f3cb867d220dfbf6fbbb2724','estudiante');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-27 16:45:17
