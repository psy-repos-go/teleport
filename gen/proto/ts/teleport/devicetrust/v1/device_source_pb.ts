/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/devicetrust/v1/device_source.proto" (package "teleport.devicetrust.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Source of device, for devices that are managed by external systems
 * (for example, MDMs).
 *
 * @generated from protobuf message teleport.devicetrust.v1.DeviceSource
 */
export interface DeviceSource {
    /**
     * Name of the source.
     * Matches the name of the corresponding MDM service, if applicable.
     * Readonly.
     *
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * Origin of the source.
     * Readonly.
     *
     * @generated from protobuf field: teleport.devicetrust.v1.DeviceOrigin origin = 2;
     */
    origin: DeviceOrigin;
}
/**
 * Origin of a device.
 *
 * @generated from protobuf enum teleport.devicetrust.v1.DeviceOrigin
 */
export enum DeviceOrigin {
    /**
     * Unspecified or absent origin.
     *
     * @generated from protobuf enum value: DEVICE_ORIGIN_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * Devices originated from direct API usage.
     *
     * @generated from protobuf enum value: DEVICE_ORIGIN_API = 1;
     */
    API = 1,
    /**
     * Devices originated from Jamf sync.
     *
     * @generated from protobuf enum value: DEVICE_ORIGIN_JAMF = 2;
     */
    JAMF = 2,
    /**
     * Source originated from Microsoft Intune sync.
     *
     * @generated from protobuf enum value: DEVICE_ORIGIN_INTUNE = 3;
     */
    INTUNE = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class DeviceSource$Type extends MessageType<DeviceSource> {
    constructor() {
        super("teleport.devicetrust.v1.DeviceSource", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "origin", kind: "enum", T: () => ["teleport.devicetrust.v1.DeviceOrigin", DeviceOrigin, "DEVICE_ORIGIN_"] }
        ]);
    }
    create(value?: PartialMessage<DeviceSource>): DeviceSource {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.name = "";
        message.origin = 0;
        if (value !== undefined)
            reflectionMergePartial<DeviceSource>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeviceSource): DeviceSource {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* teleport.devicetrust.v1.DeviceOrigin origin */ 2:
                    message.origin = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: DeviceSource, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* teleport.devicetrust.v1.DeviceOrigin origin = 2; */
        if (message.origin !== 0)
            writer.tag(2, WireType.Varint).int32(message.origin);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.devicetrust.v1.DeviceSource
 */
export const DeviceSource = new DeviceSource$Type();
