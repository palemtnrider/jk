import std from 'std';

const info = std.info('testfs/foo.txt');
std.write(info, 'fileinfo.json');

const dir = std.dir('testfs');
std.write(dir, 'dir.json');
