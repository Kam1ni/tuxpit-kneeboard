local log = require("log")
local socket = require("socket")

local LOG_KEY = "TUXPIT-KNEEBOARD"

log.write(LOG_KEY, log.INFO, "Script loaded")

function printTable(val, name, skipnewlines, depth)
    skipnewlines = skipnewlines or false
    depth = depth or 0

    local tmp = string.rep(" ", depth)
	tmp = tmp .. " " .. (name or "root").. " "

    if type(val) == "table" then
        tmp = tmp .. "{" .. (not skipnewlines and "\n" or "")

        for k, v in pairs(val) do
            tmp =  tmp .. printTable(v, k, skipnewlines, depth + 1) .. "," .. (not skipnewlines and "\n" or "")
        end

        tmp = tmp .. string.rep(" ", depth) .. "}"
    elseif type(val) == "number" then
        tmp = tmp .. tostring(val)
    elseif type(val) == "string" then
        tmp = tmp .. string.format("%q", val)
    elseif type(val) == "boolean" then
        tmp = tmp .. (val and "true" or "false")
    else
        tmp = tmp .. "\"[inserializeable datatype:" .. type(val) .. "]\""
    end

    return tmp
end

function send(toSetProperty, value)
	local udp = socket.udp()
	udp:setpeername("127.0.0.1", 8912)
	udp:send(toSetProperty .. ":" .. value)
end


local callbacks = {}

function callbacks.onMissionLoadEnd()
	log.write(LOG_KEY, log.INFO, "Mission loaded")
	sendMissionInfo()
end

function sendPlaneInfo() 
	log.write(LOG_KEY, log.INFO, "Simulation started")
	local playerId = Export.LoGetPlayerPlaneId()
	if playerId == nil then
		return
	end
	local planeName = Export.LoGetObjectById(playerId).Name
	log.write(LOG_KEY, log.INFO, "Plane type " .. planeName)
	send("aircraft", planeName)
end

function sendTerrainInfo()
	local mission = DCS.getCurrentMission()
	send("terrain", mission.mission.theatre)
end

function sendMissionInfo()
	local success, result = pcall(DCS.getMissionFilename)
	if success and result then
		log.write(LOG_KEY, log.INFO, "Mission: " .. result)
		send("mission", result)
	end
end

function callbacks.onSimulationStart()
	sendPlaneInfo()
	sendTerrainInfo()
end

function callbacks.onPlayerChangeSlot()
	log.write(LOG_KEY, log.INFO, "Player changed slots")
	sendPlaneInfo()
end

function callbacks.onPlayerStart()
	log.write(LOG_KEY, log.INFO, "On player start " .. printTable(Export))
end

--function callbacks.onSimulationFrame()
--	log.write(LOG_KEY, log.INFO, "On simulation frame " .. printTable(Export))
--end

DCS.setUserCallbacks(callbacks)
