/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

////////////////////////////////////////////////////////////////
// Fins Message Structure
////////////////////////////////////////////////////////////////

// Fins Message
[type FinsMessage (bit response)
    [const              uint 32             protocolId      0x46494E53]                         // Fins Protocol Identifier
    [implicit           uint 32             length          'messageBody.lengthInBytes + 8']                // todo: Can this calculate the length of the FinsMessage?? Packet length, automatically calculated
    [simple             uint 32             command]                                            // Handshake is 0x0000000; read and write is 0x0000002
    [const              uint 32             errorCode       0x00000000]
    [simple             FinsMessageBody('command','response')   messageBody]
]

[discriminatedType FinsMessageBody (uint 32 command, bit response)
    [typeSwitch command,response
        ['0x00000000','false' FinsHandshakeRequest                                              // Handshake request
            [const      uint 8  CLIENT      0x00]
            [const      uint 8  SNA         0x00]
            [const      uint 8  SA2         0x00]
            [simple     uint 8  SA1             ]

        ]
        ['0x00000001','false' FinsHandshakeResponse
            [const      uint 8  CLIENT      0x00]
            [const      uint 8  SNA         0x00]
            [const      uint 8  SA2         0x00]                                                      // Client node
            [simple     uint 8  SA1             ]
            [const      uint 8  SERVER      0x00]
            [const      uint 8  DNA         0x00]
            [const      uint 8  DA2         0x00]                                                       // Server node
            [simple     uint 8  DA1             ]
        ]
        ['0x00000002','false' FinsReadRequest
            [simple FinsHeaderSection section]
            [simple FinsSrcMrcCode srcMrcCode]                                                  // Command code
            [simple FinsRegisterAddress registerAddress]                                        // Register address
            [simple uint 16 startAddress]                                                       // Starting address
            [simple uint 8  bitAddress]                                                         // bit address
            [simple uint 16 dataLength]                                                         // Data length
        ]
        ['0x00000002','true' FinsReadResponse
            [simple FinsHeaderSection section]
            [simple FinsSrcMrcCode srcMrcCode]                                                  // Command code
            [simple uint 16 resErrorCode]                                                       // Error code 2 bytes
            [implicit   uint 8      byteCount     'COUNT(value)']
            [array      byte        value         count 'byteCount']                            // Returned variable data content
        ]
        ['0x00000002','false' FinsWriteRequest
            [simple FinsHeaderSection section]
            [simple FinsRegisterAddress registerAddress]                                        // Register address
            [simple uint 16 startAddress]                                                       // Starting address
            [simple uint 8  bitAddress]                                                         // bit address
            [implicit   uint 8      byteCount     'COUNT(value)']
            [array      byte        value         count   'byteCount']
        ]
        ['0x00000002','true' FinsWriteResponse
            [simple FinsHeaderSection section]
            [simple uint 16 resErrorCode]                                                       // Error code 2 bytes
        ]
    ]
]

[type ClientNode
    [const      uint 8  SA1     0x00]
    [simple     uint 8  SNA     ]
    [const      uint 8  SA2     0x00]
]

[type ServerNode
    [const      uint 8  DNA     0x00]
    [simple     uint 8  DA1     ]
    [const      uint 8  DA2     0x00]
]

[type FinsHeaderSection
    [simple uint 8 ICF]                     // ICF: 80
    [const uint 8 RSV 0x00]                 // RSV: Fixed 00
    [const uint 8 GCT 0x02]                 // GCT: Fixed 02
    [simple uint 8 DNA]                     // DNA: Destination Network Address
    [simple uint 8 DA1]                     // DA1: Destination Node Address
    [simple uint 8 DA2]                     // DA2: Destination Unit Address
    [simple uint 8 SNA]                     // SNA: Source Network Address
    [simple uint 8 SA1]                     // SA1: Source Node Address
    [simple uint 8 SA2]                     // SA2: Source Unit Address
    [const int 8 SID 0x00]                  // SID: Fixed 00
]

// Protocol command code enumeration
[enum FinsCommandCode
    ['0x00000001' FinsReadRequest]            // Read Request
    ['0x00000002' FinsWriteRequest]           // Write Request
    ['0x00000003' FinsStatusRequest]          // Status Request
    ['0x00000004' FinsDiagnosticRequest]      // Diagnostic Request
    ['0x00000005' FinsDeviceSettingRequest]   // Device Setting Request
    ['0x00000006' FinsControlRequest]         // Control Request
    ['0x00000007' FinsUserDataRequest]        // User Data Request
    ['0x00000008' FinsTerminationRequest]     // Termination Request
    ['0x00000000' FinsHandshakeRequest]       // Handshake Request
]

// Error code
[enum uint 32 FinsErrorCode
    ['0x00000000' NoError]                // No Error
    ['0x00000001' InvalidCommand]         // Invalid Command
    ['0x00000002' InvalidAddress]         // Invalid Address
    ['0x00000003' InvalidData]            // Invalid Data
    ['0x00000004' TimeoutError]           // Timeout Error
    ['0x00000005' ConnectionError]        // Connection Error
]


// Equipment status code
[enum uint 8 FinsDeviceStatus
    ['0x00' Normal]                      // Normal Status
    ['0x01' Error]                       // Error Status
    ['0x02' Busy]                        // Busy Status
    ['0x03' Offline]                     // Offline Status
]

[enum uint 8 FinsRegisterAddress
    ['0x82' DM]   // DM Address
    ['0x90' AI]   // AI Address
    ['0x91' AO]   // AO Address
    ['0xA0' HR]   // HR Address
    ['0xB0' AR]   // AR Address
]

// Main Request Code
[enum uint 8 FinsMRC
    ['0x01' Read]                       // Read Request
    ['0x02' Write]                      // Write Request
    ['0x03' Status]                     // Status Request
    ['0x04' Diagnostic]                 // Diagnostic Request
    ['0x05' DeviceSetting]              // Device Setting Request
    ['0x06' Control]                    // Control Request
    ['0x07' UserData]                   // User Data Request
    ['0x08' Termination]                // Termination Request
]

// Secondary Request Code
[enum uint 8 FinsSRC
    ['0x01' Normal]           // Normal Operation
    ['0x02' Repeat]           // Repeat Operation
    ['0x03' DataTransfer]     // Data Transfer Operation
    ['0x04' Configuration]    // Configuration Operation
    ['0x05' Force]            // Force Operation
    ['0x06' Control]          // Control Operation
]

[enum uint 16 FinsSrcMrcCode
    ['0x0101' FinsReadRequest]                      // Read Request
    ['0x0102' FinsWriteRequest]                     // Write Request
    ['0x0103' FinsStatusRequest]                    // Status Request
    ['0x0104' FinsDiagnosticRequest]                // Diagnostic Request
    ['0x0105' FinsDeviceSettingRequest]             // Device Setting Request
    ['0x0106' FinsControlRequest]                   // Control Request
    ['0x0107' FinsUserDataRequest]                  // User Data Request
    ['0x0108' FinsTerminationRequest]               // Termination Request
]


[dataIo DataItem(OmronFinsDataType dataType, uint 16 numberOfValues)
    [typeSwitch dataType,numberOfValues
        ['BOOL','1'  BOOL
            // TODO: Possibly change the order of the bit and the reserved part.
            [reserved uint 15 '0x0000'                         ]
            [simple   bit     value                            ]
        ]
        ['BOOL'      List
            // TODO: Handle adding some reserved bits at the end to fill up the last word.
            [array    bit     value count 'numberOfValues'     ]
        ]
        ['BYTE','1'  BYTE
            [reserved uint 8 '0x00']
            [simple uint 8 value]
        ]
        ['BYTE' List
            // TODO: If the number of values is odd, add a reserved byte
            [array    bit     value count 'numberOfValues * 8' ]
        ]
        ['USINT','1' USINT
            [simple uint 8 value]
        ]
        ['USINT' List
            [array uint 8 value count 'numberOfValues']
        ]
        ['SINT','1' SINT
            [simple int 8 value]
        ]
        ['SINT' List
            [array int 8 value count 'numberOfValues']
        ]
        ['UINT','2' UINT
            [simple uint 16 value]
        ]
        ['UINT' List
            [array uint 16 value count 'numberOfValues']
        ]
        ['INT','2' INT
            [simple int 16 value]
        ]
        ['INT' List
            [array int 16 value count 'numberOfValues']
        ]
        ['UDINT','4' UDINT
            [simple uint 32 value]
        ]
        ['UDINT' List
            [array uint 32 value count 'numberOfValues']
        ]
        ['DINT','4' DINT
            [simple int 32 value]
        ]
        ['DINT' List
            [array int 32 value count 'numberOfValues']
        ]
        ['ULINT','8' ULINT
            [simple uint 64 value]
        ]
        ['ULINT' List
            [array uint 64 value count 'numberOfValues']
        ]
        ['LINT','8' LINT
            [simple int 64 value]
        ]
        ['LINT' List
            [array int 64 value count 'numberOfValues']
        ]
        ['REAL','1' REAL
            [simple float 32  value]
        ]
        ['REAL' List
            [array float 32 value count 'numberOfValues']
        ]
        ['LREAL','1' LREAL
            [simple float 64  value]
        ]
        ['LREAL' List
            [array float 64 value count 'numberOfValues']
        ]
        ['CHAR','1' CHAR
            [simple string 8 value encoding='"UTF-8"']
        ]
        ['CHAR' List
            [array string 8 value count 'numberOfValues' encoding='"UTF-8"']
        ]
        ['WCHAR','1' WCHAR
            [simple string 16 value encoding='"UTF-16"']
        ]
        ['WCHAR' List
            [array string 16 value count 'numberOfValues' encoding='"UTF-16"']
        ]
    ]
]

[enum uint 8 OmronFinsDataType(uint 8 dataTypeSize)
    ['1' USINT ['1']]
    ['2' SINT ['1']]
    ['3' INT ['2']]
    ['4' UINT ['2']]
    ['5' DINT ['4']]
    ['6' UDINT ['4']]
    ['7' LINT ['8']]
    ['8' ULINT ['8']]
    ['9' REAL ['4']]
    ['10' LREAL ['8']]
    ['11' BOOL ['1']]
    ['12' BYTE ['2']]
    ['13' CHAR ['1']]
    ['14' WCHAR ['2']]
    ['15' STRING ['1']]
    ['16' WSTRING ['2']]
]