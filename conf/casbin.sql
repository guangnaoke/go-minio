-- Root
-- User
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/user/registered','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/user/login','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/user/info','GET','','','');
-- Buckets
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/buckets/list','GET','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/buckets/exists','GET','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/buckets/remove','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/buckets/listobjects','GET','','','');
-- Object
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/object/stat','GET','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/object/remove','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/object/upload','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','root','/api/object/url','GET','','','');

-- No user registered just reads and writes
-- readwrite
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/user/login','POST','','','');
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/user/info','GET','','','');
-- Buckets
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/buckets/list','GET','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/buckets/exists','GET','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/buckets/remove','POST','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/buckets/listobjects','GET','','','');
-- Object
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/object/stat','GET','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/object/remove','POST','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/object/upload','POST','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','readwrite','/api/object/url','GET','','','');

-- No user registered just reads and writes
-- read
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','read','/api/user/login','POST','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','read','/api/user/info','GET','','','');
-- Buckets
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','read','/api/buckets/list','GET','','','');
INSERT INTO casbin_rule (ptype,v0,v1,v2,v3,v4,v5) VALUES (
'p','read','/api/buckets/listobjects','GET','','','');