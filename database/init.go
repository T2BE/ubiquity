/**
 * Copyright 2017 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package database

import (
    "os"
    "github.com/IBM/ubiquity/utils/logs"
)

const KeyPsqlHost = "UBIQUITY_DB_PSQL_HOST"
const KeySqlitePath = "UBIQUITY_DB_SQLITE_PATH"

func InitPostgres(hostname string) func() {
    defer logs.GetLogger().Trace(logs.DEBUG)()
    return initConnectionFactory(&postgresFactory{host: hostname})
}

func InitSqlite(filepath string) func() {
    defer logs.GetLogger().Trace(logs.DEBUG)()
    return initConnectionFactory(&sqliteFactory{path: filepath})
}

func InitTestError() func() {
    defer logs.GetLogger().Trace(logs.DEBUG)()
    return initConnectionFactory(&testErrorFactory{})
}

func Initialize() func() {
    defer logs.GetLogger().Trace(logs.DEBUG)()

    psqlHost := os.Getenv(KeyPsqlHost)
    if psqlHost != "" {
        return InitPostgres(psqlHost)
    }
    return InitSqlite(os.Getenv(KeySqlitePath))
}
