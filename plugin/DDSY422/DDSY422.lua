package.path = "./plugin/DDSY422/?.lua;"

require "utils"
require "json"


--购电次数, (k,v) (表具地址, 购电次数)
purchaseCount = {}

-- 公共函数---------------------------------------------------------------------------------------------------------------
-- 校验和
function checkSum(buffer, startIndex, endIndex)
    local sum = 0

    for i = startIndex, endIndex do
        sum = sum + buffer[i]
    end

    return sum % 256
end

-- 地址转换 2 ~ 7 字节
function convertAddress(requestADU, sAddr)
    local addr = string.format("%012d", tonumber(sAddr))
    requestADU[2] = tonumber(string.sub(addr, 11, 12), 16)
    requestADU[3] = tonumber(string.sub(addr, 9, 10), 16)
    requestADU[4] = tonumber(string.sub(addr, 7, 8), 16)
    requestADU[5] = tonumber(string.sub(addr, 5, 6), 16)
    requestADU[6] = tonumber(string.sub(addr, 3, 4), 16)
    requestADU[7] = tonumber(string.sub(addr, 1, 2), 16)
end

-- 加密
function encode(source, startIndex, endIndex)
    local y = 13
    for i = startIndex, endIndex do
        local _, dec = math.modf(y / 2)
        if (dec == 0) then
            source[i] = (source[i] + 0x33 + 0x48 + y) % 256
        else
            source[i] = (source[i] + 0x33 + 0x54 + y) % 256
        end
        y = y + 1
    end
end

-- 解码
function decode(source, startIndex, endIndex)
    local result = {}

    local y = 1
    for i = startIndex, endIndex do
        local _, dec = math.modf(y / 2)
        if (dec == 0) then
            result[y] = source[i] - 0x33 - 0x54 - y + 1
        else
            result[y] = source[i] - 0x33 - 0x48 - y + 1
        end

        if (result[y] < 0) then
            result[y] = result[y] + 256
        end
        y = y + 1
    end

    return result
end
------------------------------------------------------------------------------------------------------------------------
-- 生成获取示数的指令
function GenerateGetConsumption(sAddr, continued)
    --                         |<-------------地址----------->|                                checkSum
    --                   1     2     3     4     5     6     7     8     9     10    11    12    13    14
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x01, 0x02, 0x8B, 0x18, 0x00, 0x16 }

    -- 2 ~ 7 字节
    convertAddress(requestADU, sAddr)

    requestADU[13] = checkSum(requestADU, 1, 12)

    return { Status = continued, Variable = requestADU }
end

-- 生成获取剩余金额、欠费金额的指令
function GenerateGetAmount(sAddr, continued)
    --                         |<-------------地址----------->|                                checkSum
    --                   1     2     3     4     5     6     7     8     9     10    11    12    13    14
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x01, 0x02, 0xBA, 0x78, 0x00, 0x16 }
    -- 2 ~ 7 字节
    convertAddress(requestADU, sAddr)

    requestADU[13] = checkSum(requestADU, 1, 12)

    return { Status = continued, Variable = requestADU }
end

-- 产生获取设备属性（遥测数据）相关指令
function GenerateGetRealVariables(sAddr, step)
    print("DDSY422 ver 1.2")
    if (step == 0) then
        return GenerateGetConsumption(sAddr, "1")
    elseif (step == 1) then
        return GenerateGetAmount(sAddr, "0")
    end
end
------------------------------------------------------------------------------------------------------------------------
function DeviceCustomCmd_SetControl_Step(sAddr, value)
    --                         |<-------------地址----------->|                    checkSum
    --                   1     2     3     4     5     6     7     8     9     10    11    12
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x16 }
    -- 2 ~ 7 字节
    convertAddress(requestADU, sAddr)

    requestADU[9] = value

    requestADU[11] = checkSum(requestADU, 1, 10)

    return { Status = "0", Variable = requestADU }
end

-- 恢复本地控制
function DeviceCustomCmd_SetControl_2(sAddr)

    return DeviceCustomCmd_SetControl_Step(sAddr, 0x16)
end
-- 强制分闸
function DeviceCustomCmd_SetControl_0(sAddr)
    return DeviceCustomCmd_SetControl_Step(sAddr, 0x15)
end

-- 强制合闸
function DeviceCustomCmd_SetControl_1(sAddr)
    return DeviceCustomCmd_SetControl_Step(sAddr, 0x14)
end

function DeviceCustomCmd_SetControl(sAddr, params)
    if (params["control"] == 0) then
        return DeviceCustomCmd_SetControl_0(sAddr)
    elseif (params["control"] == 1) then
        return DeviceCustomCmd_SetControl_1(sAddr)
    elseif (params["control"] == 2) then
        return DeviceCustomCmd_SetControl_2(sAddr)
    end

    return { Status = "0", Variable = {} }
end

-- 获取购电次数
function DeviceCustomCmd_GetCount(sAddr)
    --                   1     2     3     4     5     6     7     8     9     10    11    12    13    14
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x01, 0x02, 0xfe, 0x7c, 0x00, 0x16 }
    -- 2 ~ 7 字节
    convertAddress(requestADU, sAddr)

    requestADU[13] = checkSum(requestADU, 1, 12)

    return { Status = "0", Variable = requestADU }
end

-- 获取剩余金额
function DeviceCustomCmd_GetAmount(sAddr)
    --                         |<-------------地址----------->|                                checkSum
    --                   1     2     3     4     5     6     7     8     9     10    11    12    13    14
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x01, 0x02, 0xba, 0x78, 0x00, 0x16 }
    -- 2 ~ 7 字节
    convertAddress(requestADU, sAddr)

    requestADU[13] = checkSum(requestADU, 1, 12)

    return { Status = "0", Variable = requestADU }
end

function DeviceCustomCmd_AddAmount(sAddr, params)
    -- 68 81 24 73 00 00 00 68 04 12 0A 7C E2 CD A0 9D 92 97 D5 A1 AD 9A A8 D8 BC B8 DD 98 BF 16
    --                         |<-------------地址----------->|
    --                   1     2     3     4     5     6     7     8     9     10    11    12    13    14    15
    local requestADU = { 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x04, 0x12, 0x0A, 0x7C, 0xE2, 0xCD, 0xA0,
        --                                                                                          checkSum
        --               16    17    18    19    20    21    22    23    24    25    26    27    28    29    30
                         0x9D, 0x92, 0x97, 0xD5, 0xA1, 0xAD, 0x9A, 0xA8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16 }
    -- 地址转换
    convertAddress(requestADU, sAddr)

    -- 充值金额，遗留问题：退费
    local charge = math.abs(params["amount"]) * 100
    local amount = string.format("%06d", charge)
    requestADU[24] = tonumber(string.sub(amount, 5, 6), 16)
    requestADU[25] = tonumber(string.sub(amount, 3, 4), 16)
    requestADU[26] = tonumber(string.sub(amount, 1, 2), 16)

    -- 购电次数
    if purchaseCount[sAddr] == nil then
        purchaseCount[sAddr] = 1
    end

    purchaseCount[sAddr] = purchaseCount[sAddr] % 10000
    local buy_times = string.format("%04d", purchaseCount[sAddr] + 1)
    requestADU[27] = tonumber(string.sub(buy_times, 3, 4), 16)
    requestADU[28] = tonumber(string.sub(buy_times, 1, 2), 16)

    -- 加密
    encode(requestADU, 24, 28)

    requestADU[29] = checkSum(requestADU, 1, 28)

    return { Status = "0", Variable = requestADU }
end

function DeviceCustomCmd(sAddr, cmdName, cmdParam, step)
    local params = json.jsondecode(cmdParam)

    print(sAddr, cmdName, cmdParam, step)

    if (cmdName == "SetControl") then
        return DeviceCustomCmd_SetControl(sAddr, params)
    elseif (cmdName == "GetCount") then
        return DeviceCustomCmd_GetCount(sAddr)
    elseif (cmdName == "GetAmount") then
        return DeviceCustomCmd_GetAmount(sAddr)
    elseif (cmdName == "AddAmount") then
        return DeviceCustomCmd_AddAmount(sAddr, params)
    end

    -- 出错了，没有对应的函数名
    return { Status = "0", Variable = {} }
end

------------------------------------------------------------------------------------------------------------------------
-- go 程序接收到表具数据，放到rxBuf中，从而供AnalysisRx解析
rxBuf = {}

function AnalysisRx(sAddr, rxBufCnt)
    -- 最短指令12个字节
    if (rxBufCnt < 12) then
        rxBuf = {}
        -- Status = "1" 错误
        return { Status = "1", Variable = {} }
    end

    -- 忽略掉头部不正确的字节数据
    for index = 1, rxBufCnt do
        repeat
            -- 如果剩余数据不足12字节，则退出
            if (index + 11 > rxBufCnt) then
                rxBuf = {}
                -- Status = "1" 错误
                return { Status = "1", Variable = {} }
            end
            -- 地址（2~7字节）两端不是0x68，数据格式错误，则退出
            if ((rxBuf[index] ~= 0x68) or (rxBuf[index + 7] ~= 0x68)) then
                break
            end
            -----------------------------------------------
            -- 地址不对，则退出
            local addr = string.format("%012d", tonumber(sAddr))
            local mAddr = string.format("%02X%02X%02X%02X%02X%02X",
                    rxBuf[index + 6], rxBuf[index + 5], rxBuf[index + 4], rxBuf[index + 3], rxBuf[index + 2], rxBuf[index + 1])
            if (mAddr ~= addr) then
                print("error:mAddr ~= addr")
                rxBuf = {}
                -- Status = "1" 错误
                return { Status = "1", Variable = {} }
            end

            -----------------------------------------------
            local funCode = rxBuf[index + 8]
            local dataLen = rxBuf[index + 9]
            -- 数据长度不足，则退出
            if (index + 10 + dataLen > rxBufCnt) then
                rxBuf = {}
                -- Status = "1" 错误
                return { Status = "1", Variable = {} }
            end
            ----------------------------------------------
            -- 累加和不正确，则退出
            if (checkSum(rxBuf, index, index + 9 + dataLen) ~= rxBuf[index + 10 + dataLen]) then
                rxBuf = {}
                -- Status = "1" 错误
                return { Status = "1", Variable = {} }
            end
            ----------------------------------------------
            -- 以下是正确数据
            if (funCode == 0x81) then
                -- 0x81 正确应答
                ------------------------------------------
                --取出数据域
                local pdu = decode(rxBuf, index + 10, index + 10 + dataLen)
                local mDataAddr = pdu[1] * 256 + pdu[2]
                if (mDataAddr == 0x1090) then
                    local dev_consumption = (math.modf(pdu[6] / 16) * 10 + math.modf(pdu[6] % 16)) * 1000000 +
                            (math.modf(pdu[5] / 16) * 10 + math.modf(pdu[5] % 16)) * 10000 +
                            (math.modf(pdu[4] / 16) * 10 + math.modf(pdu[4] % 16)) * 100 +
                            math.modf(pdu[3] / 16) * 10 + math.modf(pdu[3] % 16)

                    -- 清除缓存
                    rxBuf = {}
                    return {
                        Status = "0",
                        Variable = {
                            utils.AppendVariable(0, "dev_consumption", "总电能", "double", dev_consumption,
                                    string.format("%3.2f", dev_consumption / 100.0))
                        }
                    }
                elseif (mDataAddr == 0x3FF0) then
                    -- 0x3FF0 欠费剩余金额标识符
                    local dev_remain_amt = (math.modf(pdu[6] / 16) * 10 + math.modf(pdu[6] % 16)) * 1000000 +
                            (math.modf(pdu[5] / 16) * 10 + math.modf(pdu[5] % 16)) * 10000 +
                            (math.modf(pdu[4] / 16) * 10 + math.modf(pdu[4] % 16)) * 100 +
                            math.modf(pdu[3] / 16) * 10 + math.modf(pdu[3] % 16)

                    local dev_owe_amt = (math.modf(pdu[10] / 16) * 10 + math.modf(pdu[10] % 16)) * 1000000 +
                            (math.modf(pdu[9] / 16) * 10 + math.modf(pdu[9] % 16)) * 10000 +
                            (math.modf(pdu[8] / 16) * 10 + math.modf(pdu[8] % 16)) * 100 +
                            math.modf(pdu[7] / 16) * 10 + math.modf(pdu[7] % 16)

                    -- 清除缓存
                    rxBuf = {}
                    return {
                        Status = "0",
                        Variable = {
                            utils.AppendVariable(1, "dev_remain_amt", "剩余金额", "double", dev_remain_amt,
                                    string.format("%3.2f", dev_remain_amt / 100.0)),
                            utils.AppendVariable(2, "dev_owe_amt", "欠费金额", "double", dev_owe_amt,
                                    string.format("%3.2f", dev_owe_amt / 100.0))
                        }
                    }
                elseif (mDataAddr == 0x83F4) then
                    -- 0x83F4 购电次数标识符
                    purchaseCount[sAddr] = (math.modf(pdu[4] / 16) * 10 + math.modf(pdu[4] % 16)) * 100 +
                            math.modf(pdu[3] / 16) * 10 + math.modf(pdu[3] % 16)
                    -- 正确
                    rxBuf = {}
                    return { Status = "0", Variable = {} }
                end
                ------------------------------------------
            elseif ((funCode == 0x14) or
                    (funCode == 0x15) or
                    (funCode == 0x84)) then
                -- 0x14 合闸成功
                -- 0x15 分闸成功
                -- 0x84 充值成功
                -- 正确
                rxBuf = {}
                return { Status = "0", Variable = {} }
            end
            break
        until true
    end

    -- 错误
    rxBuf = {}
    return { Status = "1", Variable = {} }
end

------------------------------------------------------------------------------------------------------------------------
-- 以下是单元测试
--[[
白向阳2：测试完成，用多行注释形式，将以下程序注释掉

]]--
function TestGenerateGetRealVariablesStep(step)
    print("GenerateGetRealVariables：step=0")
    local a = GenerateGetRealVariables("732481", step)
    print("status=", a.Status)

    local b = ""
    for _, v in ipairs(a.Variable) do
        b = b .. string.format("%02x ", v)
    end
    print("variable=", b)

end

function TestGenerateGetRealVariables()
    -- 68 64 00 00 00 00 00 68 01 02 8B 18 DA 16
    TestGenerateGetRealVariablesStep(0)
    -- 68 64 00 00 00 00 00 68 01 02 BA 78 69 16
    TestGenerateGetRealVariablesStep(1)
end

function TestDeviceCustomCmd_Params(cmdName, cmdParams)
    print("函数名:", cmdName)

    local a = DeviceCustomCmd("732481", cmdName, cmdParams, 0)
    print("status = ", a.Status)

    local b = ""
    for _, v in ipairs(a.Variable) do
        b = b .. string.format("%02x ", v)
    end
    print("variable = ", b)
end

function TestDeviceCustomCmd()
    -- 68 12 00 00 00 00 00 68 15 00 f7 16
    TestDeviceCustomCmd_Params("SetControl", "{\"control\": 0}")
    -- 68 12 00 00 00 00 00 68 14 00 f6 16
    TestDeviceCustomCmd_Params("SetControl", "{\"control\": 1}")
    --                    我猜是13或16   ？？？？？？？？？？？？？？？？？？？？？？？
    TestDeviceCustomCmd_Params("SetControl", "{\"control\": 2}")
    -- 68 12 00 00 00 00 00 68 01 02 fe 7c 275 16
    TestDeviceCustomCmd_Params("GetCount", "{}")
    -- 68 12 00 00 00 00 00 68 01 02 ba 78 22d 16
    TestDeviceCustomCmd_Params("GetAmount", "{}")
    purchaseCount["732481"] = 22
    --68 12 00 00 00 00 00 68 04 12 0a 7c e2 cd a0 9d 92 97 d5 a1 ad 98 a8 94 89 97 ae 98 bf0 16
    TestDeviceCustomCmd_Params("AddAmount", "{\"amount\": 100.00}")
end

function TestAnalysisRx_Consumption()
    --电能 6652.84
    rxBuf = { 0x68, 0x81, 0x24, 0x73, 0x00, 0x00, 0x00, 0x68, 0x81, 0x06, 0x8B, 0x18, 0x01, 0xDC, 0xE5, 0x8C, 0x60, 0x16 }
    local a = AnalysisRx("732481", #rxBuf)

    print("status = ", a.Status)
    if (a.Status == "0") then
        print("variable = ", a.Variable[1].Label, a.Variable[1].Name, a.Variable[1].Value)
    end
end

function TestAnalysisRx_GetCount()
    --购电次数 50
    rxBuf = { 0x68, 0x81, 0x24, 0x73, 0x00, 0x00, 0x00, 0x68, 0x81, 0x04, 0xFE, 0x7C, 0xCD, 0x8A, 0x3E, 0x16 }
    local a = AnalysisRx("732481", #rxBuf)
    print("status = ", a.Status)
    if (a.Status == "0") then
        print("purchaseCount = ", purchaseCount["732481"])
    end
end

function TestAnalysisRx()
    TestAnalysisRx_Consumption()
    TestAnalysisRx_GetCount()
end

function TestAll()
    TestGenerateGetRealVariables()

    TestDeviceCustomCmd()

    TestAnalysisRx()
end

