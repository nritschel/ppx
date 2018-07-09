<?php
// automatically generated by the FlatBuffers compiler, do not modify

namespace ppx;

use \Google\FlatBuffers\Struct;
use \Google\FlatBuffers\Table;
use \Google\FlatBuffers\ByteBuffer;
use \Google\FlatBuffers\FlatBufferBuilder;

class Normal extends Table
{
    /**
     * @param ByteBuffer $bb
     * @return Normal
     */
    public static function getRootAsNormal(ByteBuffer $bb)
    {
        $obj = new Normal();
        return ($obj->init($bb->getInt($bb->getPosition()) + $bb->getPosition(), $bb));
    }

    public static function NormalIdentifier()
    {
        return "PPXF";
    }

    public static function NormalBufferHasIdentifier(ByteBuffer $buf)
    {
        return self::__has_identifier($buf, self::NormalIdentifier());
    }

    /**
     * @param int $_i offset
     * @param ByteBuffer $_bb
     * @return Normal
     **/
    public function init($_i, ByteBuffer $_bb)
    {
        $this->bb_pos = $_i;
        $this->bb = $_bb;
        return $this;
    }

    public function getMean()
    {
        $obj = new Tensor();
        $o = $this->__offset(4);
        return $o != 0 ? $obj->init($this->__indirect($o + $this->bb_pos), $this->bb) : 0;
    }

    public function getStddev()
    {
        $obj = new Tensor();
        $o = $this->__offset(6);
        return $o != 0 ? $obj->init($this->__indirect($o + $this->bb_pos), $this->bb) : 0;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return void
     */
    public static function startNormal(FlatBufferBuilder $builder)
    {
        $builder->StartObject(2);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return Normal
     */
    public static function createNormal(FlatBufferBuilder $builder, $mean, $stddev)
    {
        $builder->startObject(2);
        self::addMean($builder, $mean);
        self::addStddev($builder, $stddev);
        $o = $builder->endObject();
        return $o;
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param int
     * @return void
     */
    public static function addMean(FlatBufferBuilder $builder, $mean)
    {
        $builder->addOffsetX(0, $mean, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @param int
     * @return void
     */
    public static function addStddev(FlatBufferBuilder $builder, $stddev)
    {
        $builder->addOffsetX(1, $stddev, 0);
    }

    /**
     * @param FlatBufferBuilder $builder
     * @return int table offset
     */
    public static function endNormal(FlatBufferBuilder $builder)
    {
        $o = $builder->endObject();
        return $o;
    }
}
