package.path = "./plugin/DTSD1352/?.lua;"

require "utils"
require "json"

regAddr = 0x0000

------------------------------------------------------------------------------------------------------------------------
-- 生成获取示数的指令
function GenerateGetConsumption(sAddr, continued)
    --                  地址                                   |checkSum|
    --                                1      2     3     4     5     6     7     8
    print("sAddr", sAddr)
    local requestADU = { tonumber(sAddr), 0x03, 0x00, 0x00, 0x00, 0x02 }
    local crc
    local VariableMapTemp = {
        Variable = requestADU
    }
    crc = GetCRCModbus(VariableMapTemp)
    requestADU[7] = math.modf(crc % 256)
    requestADU[8] = math.modf(crc / 256)
    regAddr = 0x0000
    return { Status = continued, Variable = requestADU }
end

-- 产生获取设备属性（遥测数据）相关指令
function GenerateGetRealVariables(sAddr, step)
    print("DTSD1352 ver 1.0")
    if (step == 0)
    then
        return GenerateGetConsumption(sAddr, "0")
    end
end
------------------------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------------------------------------------------------
-- go 程序接收到表具数据，放到rxBuf中，从而供AnalysisRx解析
rxBuf = {}

function AnalysisRx(sAddr, rxBufCnt)

    local index = 1
    local mAddr = 0
    local mFunCode = 0
    local mRegByte = 0
    local crcH, crcL, crc, crcTemp

    -- body
    local VariableMap = {
        Status = "1",
        Variable = {},
    }

    --modbus03 06 最短为8个字节
    if (rxBufCnt < 8)
    then
        rxBuf = {}
        return VariableMap
    end

    for index = 1, rxBufCnt, 1 do
        if (index <= rxBufCnt)
        then
            mAddr = rxBuf[index]
            mFunCode = rxBuf[index + 1]
            mNumberLen = rxBuf[index + 2]

            if (mAddr == tonumber(sAddr) and (mFunCode == 0x03) and mNumberLen == 0x04)
            then
                mRegByte = rxBuf[index + 2]
                if (mRegByte + 5 > rxBufCnt)
                then
                    rxBuf = {}
                    return VariableMap
                end
                --取出数据域数据
                local pdu = {}
                for i = index, index + 2 + mRegByte, 1 do
                    if rxBuf[i] ~= nil then
                        table.insert(pdu, rxBuf[i])
                    end
                end
                local VariableMapCRC = {
                    Variable = pdu
                }
                crc = GetCRCModbus(VariableMapCRC)
                crcTemp = rxBuf[index + 3 + mRegByte + 1] * 256 + rxBuf[index + 3 + mRegByte]
                if (crc == crcTemp)
                then
                    local Variable = {}
                    if (regAddr == 0x0000)
                    then
                        ---------------------------电能-------------------------------
                        dev_consumption = 256 * 256 * 256 * rxBuf[index + 3] + 256 * 256 * rxBuf[index + 4] + 256 * rxBuf[index + 5] + rxBuf[index + 6]
                        Variable[1] = utils.AppendVariable(0, "dev_consumption", "总电能", "double", dev_consumption, dev_consumption)
                        VariableMap = {
                            Status = "0",
                            Variable = Variable
                        }
                        rxBuf = {}
                        return VariableMap
                    end
                end
            end
        else
            rxBuf = {}
            return VariableMap
        end
    end

    rxBuf = {}
    return VariableMap
end

------------------------------------------------------------------------------------------------------------------------
-- 以下是单元测试
--[[
白向阳2：测试完成，用多行注释形式，将以下程序注释掉

]]--


function TestGenerateGetRealVariablesStep(step)
    print("GenerateGetRealVariables：step=0")
    local a = GenerateGetRealVariables("1", step)
    print("status=", a.Status)

    local b = ""
    for _, v in ipairs(a.Variable) do
        b = b .. string.format("%02x ", v)
    end
    print("variable=", b)

end

function TestGenerateGetRealVariables()
    TestGenerateGetRealVariablesStep(0)
end

function TestDeviceCustomCmd_Params(cmdName, cmdParams)
    print("函数名:", cmdName)

    local a = DeviceCustomCmd("3", cmdName, cmdParams, 0)
    print("status = ", a.Status)

    local b = ""
    for _, v in ipairs(a.Variable) do
        b = b .. string.format("%02x ", v)
    end
    print("variable = ", b)
end

--function TestDeviceCustomCmd()
--    -- 68 12 00 00 00 00 00 68 01 02 fe 7c 275 16
--    TestDeviceCustomCmd_Params("GetCount", "{}")
--    -- 68 12 00 00 00 00 00 68 01 02 ba 78 22d 16
--    TestDeviceCustomCmd_Params("GetAmount", "{}")
--    purchaseCount["3"] = 19
--    --
--    TestDeviceCustomCmd_Params("AddAmount", "{\"amount\": 10.00}")
--end

function TestAnalysisRx_Consumption()
    --电能 1820.65
    regAddr = 0x0000
    --rxBuf = { 0x03, 0x03, 0x04, 0x00, 0x00, 0x00, 0xC1, 0x18, 0x63 }
    --rxBuf = { 0x01, 0x03, 0x04, 0x00, 0x00, 0xBF, 0xF9, 0x4A, 0x41 }
    --rxBuf = { 0x01, 0x03, 0x04, 0x00, 0x01, 0x00, 0x00, 0xD8, 0xFA, 0x69, 0x01, 0x03, 0x04, 0x00, 0x00, 0xBF, 0xF9, 0x4A, 0x41 }

    --rxBuf = {0x01, 0x03, 0x04, 0x00, 0x16, 0x00, 0x39, 0xDB, 0xE5}

    rxBuf = { 0x01, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00, 0xFA, 0x33 }
    local a = AnalysisRx("1", #rxBuf)

    print("status = ", a.Status)
    if (a.Status == "0") then
        print("variable = ", a.Variable[1].Label, a.Variable[1].Name, a.Variable[1].Value)
    end
end

function TestAnalysisRx()
    TestAnalysisRx_Consumption()
end

function TestAll()
    TestGenerateGetRealVariables()

    TestAnalysisRx()
end

