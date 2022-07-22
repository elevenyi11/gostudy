from os import remove
import os.path


def readTwFile():
    fyCodeDict = {}
    folder = os.path.abspath('.')
    fyfile = os.path.join(folder, 'code.txt')
    with open(fyfile, 'r', encoding="utf-8") as f:
        for line in f.readlines():
            temps = line.strip().split('->')
            if(line.strip() == ''):
                continue

            fyCodeDict[temps[0].strip()] = temps[1].strip()

    return fyCodeDict


def readRsFile(fyCodeDict):
    zhlines = []
    folder = os.path.abspath('.')
    csfile = os.path.join(folder, 'ResultStatus.proto')
    with open(csfile, 'r', encoding="utf-8") as f:
        needConvert = False
        listLines = f.readlines()
        for i in range(len(listLines)):
            curLine = listLines[i]
            if curLine.find('enum ResultStatus {') > -1:
                needConvert = True

            if needConvert == False:
                zhlines.append(curLine)
                continue

            beforeLine = listLines[i-1]
            removeBeforeLine, newBeforeLine, newCurLine = replaceLine(
                fyCodeDict, beforeLine, curLine)

            if removeBeforeLine:
                del zhlines[i-1]
                zhlines.append(newBeforeLine)

            zhlines.append(newCurLine)

    return zhlines


def replaceLine(fyCodeDict, beforeLine, curLine):
    # 如果未找到，则认为不需要替换，直接加入结果
    index = curLine.find(' = ')
    if index == -1:
        return False, "", curLine

    # 找到key
    kLeftIndex = curLine.find(' = ')+3
    key = curLine[kLeftIndex:len(curLine)-2]
    if key not in fyCodeDict.keys():
        print('翻译文件中未找到key={}对应的翻译文档'.format(key))
        return False, '', curLine

    # 找到，该行需要替换
    rbeforeLine = "  // {}\n".format(fyCodeDict[key])

    return True, rbeforeLine, curLine


def outFile(allLine):
    folder = os.path.abspath('.')
    newCsfile = os.path.join(folder, 'ResultStatus_new.proto')
    with open(newCsfile, 'w', encoding="utf-8") as f:
        f.writelines(allLine)


# 正式开始
fyCodeDict = readTwFile()
allLine = readRsFile(fyCodeDict)
outFile(allLine)
